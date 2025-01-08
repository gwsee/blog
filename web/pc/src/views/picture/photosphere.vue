<template>
  <div ref="container" class="w-full h-screen bg-transparent">
    <!-- Loading indicator -->
    <div v-if="loading" class="absolute inset-0 flex items-center justify-center text-white">
      Loading photos...
    </div>

    <!-- Photo detail modal -->
    <a-modal
        v-model:visible="modalVisible"
        :footer="null"
        :closable="true"
        @cancel="closeModal"
        width="800px"
        class="photo-modal"
    >
      <div class="flex flex-col items-center space-y-4">
        <img
            :src="selectedPhoto?.url"
            :alt="selectedPhoto?.title"
            class="w-full h-auto rounded-lg shadow-xl max-h-[60vh] object-contain"
        />
        <h3 class="text-xl font-bold mt-4">{{ selectedPhoto?.title }}</h3>
        <p class="text-gray-600">{{ selectedPhoto?.description }}</p>
      </div>
    </a-modal>

    <!-- Hover info card -->
    <div
        v-if="hoveredPhoto"
        class="fixed px-4 py-2 bg-black/80 text-white rounded-lg text-sm pointer-events-none"
        :style="hoverCardStyle"
    >
      <p class="font-medium">{{ hoveredPhoto.title }}</p>
    </div>
  </div>
</template>

<script setup>
import {ref, onMounted, onUnmounted, onBeforeUnmount, watch} from 'vue';
import * as THREE from 'three';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls';

const props = defineProps({
  photos: {
    type: Array,
    default: () => [],
    validator: (value) => value.length <= 100
  }
});

const container = ref(null);
const modalVisible = ref(false);
const selectedPhoto = ref(null);
const hoveredPhoto = ref(null);
const loading = ref(true);
const hoverCardStyle = ref({
  display: 'none',
  left: '0px',
  top: '0px'
});

let scene, camera, renderer, controls;
let photoGroup;
let isRotating = true;
let raycaster;
let mouse;
let photoMeshes = [];

const init = () => {
  scene = new THREE.Scene();

  camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
  camera.position.z = 5;

  renderer = new THREE.WebGLRenderer({
    antialias: true,
    alpha: true,
    powerPreference: "high-performance"
  });
  renderer.setSize(window.innerWidth, window.innerHeight);
  renderer.setClearColor(0x000000, 0); // Set background to fully transparent
  container.value.appendChild(renderer.domElement);

  controls = new OrbitControls(camera, renderer.domElement);
  controls.enableDamping = true;
  controls.dampingFactor = 0.05;
  controls.rotateSpeed = 0.5;

  raycaster = new THREE.Raycaster();
  mouse = new THREE.Vector2();

  createPhotoSphere();
  loading.value = false;
};

const textureLoader = new THREE.TextureLoader();
const loadTexture = (url) => {
  return new Promise((resolve) => {
    textureLoader.load(url, (texture) => {
      console.log(texture);
      texture.minFilter = THREE.LinearFilter;
      texture.generateMipmaps = false;
      resolve(texture);
    });
  });
};
const loadTextureWithReferer = async (url) => {
  const response = await fetch(url, {
    method: 'GET',
    headers: {
      'Referer': window.location.origin, // 设置 Referer 为当前域名
    },
    redirect: 'follow', // 确保处理重定向
  });

  if (!response.ok) {
    throw new Error(`Failed to load texture: ${response.status}`);
  }

  const blob = await response.blob();
  const imageUrl = URL.createObjectURL(blob);

  // 使用 TextureLoader 加载从 blob 创建的图片
  return new Promise((resolve, reject) => {
    textureLoader.load(
        imageUrl,
        (texture) => {
          texture.minFilter = THREE.LinearFilter;
          texture.generateMipmaps = false;
          resolve(texture);
        },
        undefined,
        (err) => reject(err)
    );
  });
};
// 监听 props 的变化
watch(
    () => props.photos, // 监听 photos
    (newPhotos, oldPhotos) => {
      // 触发的方法
      createPhotoSphere()
    },
    { deep: true } // 深度监听，确保数组内容变化时也能触发
);

const createPhotoMesh = async (photo, phi, theta) => {
  const texture = await loadTexture(photo.url);
  const photoGeometry = new THREE.CircleGeometry(0.2, 16);
  const photoMaterial = new THREE.MeshBasicMaterial({
    map: texture,
    side: THREE.DoubleSide,
    transparent: true
  });

  const photoMesh = new THREE.Mesh(photoGeometry, photoMaterial);
  const radius = 2;
  photoMesh.position.setFromSpherical(new THREE.Spherical(radius, phi, theta));
  photoMesh.lookAt(0, 0, 0);
  photoMesh.userData = { photo, originalScale: photoMesh.scale.clone() };

  return photoMesh;
};

const createPhotoSphere = async () => {
  photoGroup = new THREE.Group();
  scene.add(photoGroup);
  console.log(props.photos.length)
  for (let i = 0; i < props.photos.length; i++) {
    const photo = props.photos[i];
    const phi = Math.acos(-1 + (2 * i) / props.photos.length);
    const theta = Math.sqrt(props.photos.length * Math.PI) * phi;

    const photoMesh = await createPhotoMesh(photo, phi, theta);
    photoGroup.add(photoMesh);
    photoMeshes.push(photoMesh);
  }
};

const animate = () => {
  animationFrameId = requestAnimationFrame(animate);

  if (isRotating) {
    photoGroup.rotation.y += 0.001;
  }

  updatePhotoVisibility();

  controls.update();
  checkIntersection();
  renderer.render(scene, camera);
};

const updatePhotoVisibility = () => {
  const cameraPosition = new THREE.Vector3();
  camera.getWorldPosition(cameraPosition);

  photoMeshes.forEach(mesh => {
    const distance = mesh.position.distanceTo(cameraPosition);
    if (distance > 10) {
      mesh.visible = false;
    } else {
      mesh.visible = true;
      const scale = 1 - (distance / 4) * 0.5;
      mesh.scale.set(scale, scale, scale);
    }
  });
};

const checkIntersection = () => {
  raycaster.setFromCamera(mouse, camera);
  const intersects = raycaster.intersectObjects(photoMeshes);

  // Reset all photos to original scale
  photoMeshes.forEach(mesh => {
    if (mesh.visible) {
      mesh.scale.copy(mesh.userData.originalScale);
    }
  });

  if (intersects.length > 0) {
    const intersected = intersects[0].object;
    intersected.scale.multiplyScalar(1.2);

    if (!hoveredPhoto.value || hoveredPhoto.value !== intersected.userData.photo) {
      hoveredPhoto.value = intersected.userData.photo;
      updateHoverCard(intersected);
    }
  } else {
    hoveredPhoto.value = null;
    hoverCardStyle.value.display = 'none';
  }
};

const updateHoverCard = (intersected) => {
  const vector = intersected.position.clone();
  vector.project(camera);

  const x = (vector.x * 0.5 + 0.5) * window.innerWidth;
  const y = -(vector.y * 0.5 - 0.5) * window.innerHeight;

  hoverCardStyle.value = {
    display: 'block',
    left: `${x + 10}px`,
    top: `${y - 10}px`,
    transform: 'translate(-50%, -50%)'
    //transform: 'translate(0,0)'
  };
};

const onMouseMove = (event) => {
  mouse.x = (event.clientX / window.innerWidth) * 2 - 1;
  mouse.y = -(event.clientY / window.innerHeight) * 2 + 1;

  // Check if mouse is over the sphere
  raycaster.setFromCamera(mouse, camera);
  const intersects = raycaster.intersectObjects(photoMeshes);
  isRotating = intersects.length === 0;
};

const onClick = (event) => {
  raycaster.setFromCamera(mouse, camera);
  const intersects = raycaster.intersectObjects(photoMeshes);

  if (intersects.length > 0) {
    const clicked = intersects[0].object;
    selectedPhoto.value = clicked.userData.photo;
    modalVisible.value = true;
  }
};

const closeModal = () => {
  modalVisible.value = false;
  selectedPhoto.value = null;
};

const onWindowResize = () => {
  camera.aspect = window.innerWidth / window.innerHeight;
  camera.updateProjectionMatrix();
  renderer.setSize(window.innerWidth, window.innerHeight);
};

onMounted(() => {
  init();
  animate();
  window.addEventListener('resize', onWindowResize);
  renderer.domElement.addEventListener('mousemove', onMouseMove);
  renderer.domElement.addEventListener('click', onClick);
});

onUnmounted(() => {
  window.removeEventListener('resize', onWindowResize);
  renderer.domElement.removeEventListener('mousemove', onMouseMove);
  renderer.domElement.removeEventListener('click', onClick);
  scene.traverse(object => {
    if (object.geometry) {
      object.geometry.dispose();
    }
    if (object.material) {
      if (Array.isArray(object.material)) {
        object.material.forEach(material => material.dispose());
      } else {
        object.material.dispose();
      }
    }
  });
  renderer.dispose();
});

onBeforeUnmount(() => {
  cancelAnimationFrame(animationFrameId);
});

let animationFrameId;
</script>

<style scoped>
.photo-modal :deep(.ant-modal-content) {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
}
</style>

