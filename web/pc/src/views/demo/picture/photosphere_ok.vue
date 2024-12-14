<template>
  <div ref="container" class="w-full h-screen bg-transparent">
    <!-- Loading indicator -->
    <div v-if="loading" class="absolute inset-0 flex items-center justify-center text-white">
      Loading photos...
    </div>

    <!-- Photo detail modal -->
    <a-modal
        v-model:open="modalVisible"
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
import { ref, onMounted, onUnmounted } from 'vue';
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

const createPhotoSphere = () => {
  photoGroup = new THREE.Group();
  scene.add(photoGroup);

  props.photos.forEach((photo, index) => {
    const phi = Math.acos(-1 + (2 * index) / props.photos.length);
    const theta = Math.sqrt(props.photos.length * Math.PI) * phi;

    // Create circular photo with glow effect
    const photoGeometry = new THREE.CircleGeometry(0.2, 32);
    const photoTexture = new THREE.TextureLoader().load(photo.url);
    const photoMaterial = new THREE.MeshBasicMaterial({
      map: photoTexture,
      side: THREE.DoubleSide,
      transparent: true
    });

    const photoMesh = new THREE.Mesh(photoGeometry, photoMaterial);
    const radius = 2;
    photoMesh.position.setFromSpherical(new THREE.Spherical(radius, phi, theta));
    photoMesh.lookAt(0, 0, 0);
    photoMesh.userData = { photo, originalScale: photoMesh.scale.clone() };

    // Add glow effect
    const glowGeometry = new THREE.CircleGeometry(0.25, 32);
    const glowMaterial = new THREE.MeshBasicMaterial({
      color: 0xffffff,
      transparent: true,
      opacity: 0.2,
      side: THREE.DoubleSide
    });
    const glowMesh = new THREE.Mesh(glowGeometry, glowMaterial);
    glowMesh.position.copy(photoMesh.position);
    glowMesh.lookAt(0, 0, 0);

    photoGroup.add(glowMesh);
    photoGroup.add(photoMesh);
    photoMeshes.push(photoMesh);
  });
};

const animate = () => {
  requestAnimationFrame(animate);

  if (isRotating) {
    photoGroup.rotation.y += 0.001;
  }

  controls.update();
  checkIntersection();
  renderer.render(scene, camera);
};

const checkIntersection = () => {
  raycaster.setFromCamera(mouse, camera);
  const intersects = raycaster.intersectObjects(photoMeshes);

  // Reset all photos to original scale
  photoMeshes.forEach(mesh => {
    mesh.scale.copy(mesh.userData.originalScale);
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
    left: `${x}px`,
    top: `${y - 50}px`,
    transform: 'translate(-50%, -50%)'
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
  renderer.dispose();
});
</script>

<style scoped>
.photo-modal :deep(.ant-modal-content) {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
}
</style>

