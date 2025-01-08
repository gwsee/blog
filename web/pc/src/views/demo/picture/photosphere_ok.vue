<template>
  <div ref="viewer" class="w-full h-full">
    <div
        v-if="hoveredPhoto"
        class="fixed px-4 py-2 bg-black/80 text-white rounded-lg text-sm pointer-events-none transition-all duration-200 ease-out"
        :style="hoverCardStyle"
    >
      {{ hoveredPhoto.title }}
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, watch } from 'vue';
import * as THREE from 'three';

const viewer = ref(null);
let scene, camera, renderer, raycaster, mouse;
let intersected = null;
const hoveredPhoto = ref(null);
const hoverCardStyle = reactive({
  display: 'none',
  left: 0,
  top: 0,
});

onMounted(() => {
  init();
  animate();
});

const init = () => {
  scene = new THREE.Scene();
  camera = new THREE.PerspectiveCamera(75, viewer.value.offsetWidth / viewer.value.offsetHeight, 0.1, 1000);
  renderer = new THREE.WebGLRenderer({ antialias: true });
  renderer.setSize(viewer.value.offsetWidth, viewer.value.offsetHeight);
  viewer.value.appendChild(renderer.domElement);

  raycaster = new THREE.Raycaster();
  mouse = new THREE.Vector2();

  // Add your photo objects here...  For example:
  const geometry = new THREE.BoxGeometry(1, 1, 1);
  const material = new THREE.MeshBasicMaterial({ color: 0x00ff00 });
  const cube = new THREE.Mesh(geometry, material);
  cube.position.set(0, 0, -5);
  cube.userData = { title: 'Photo 1' }; // Add data to identify the photo
  scene.add(cube);


  camera.position.z = 5;
};

const animate = () => {
  requestAnimationFrame(animate);
  render();
};

const render = () => {
  raycaster.setFromCamera(mouse, camera);
  const intersects = raycaster.intersectObjects(scene.children);

  if (intersects.length > 0) {
    if (intersected !== intersects[0].object) {
      intersected = intersects[0].object;
      updateHoverCard(intersected);
      hoveredPhoto.value = intersected.userData;
    }
  } else {
    hoveredPhoto.value = null;
    hoverCardStyle.value.display = 'none';
    intersected = null;
  }
  renderer.render(scene, camera);
};

const updateHoverCard = (intersected) => {
  const vector = intersected.position.clone();
  vector.project(camera);

  const x = (vector.x * 0.5 + 0.5) * window.innerWidth;
  const y = -(vector.y * 0.5 - 0.5) * window.innerHeight;

  // Calculate offset for top-right positioning
  const offset = 10; // 10px offset from the element

  hoverCardStyle.value = {
    display: 'block',
    left: `${x + offset}px`,
    top: `${y - offset}px`,
    transform: 'translate(0, 0)' // Remove the -50% translation to keep it anchored
  };
};


window.addEventListener('mousemove', (event) => {
  mouse.x = (event.clientX / window.innerWidth) * 2 - 1;
  mouse.y = -(event.clientY / window.innerHeight) * 2 + 1;
});

window.addEventListener('resize', () => {
  camera.aspect = viewer.value.offsetWidth / viewer.value.offsetHeight;
  camera.updateProjectionMatrix();
  renderer.setSize(viewer.value.offsetWidth, viewer.value.offsetHeight);
});
</script>

<style scoped>
.fixed {
  position: fixed;
}
</style>
