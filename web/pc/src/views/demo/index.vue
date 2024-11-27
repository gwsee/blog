<template>
  <div ref="container" class="w-full h-screen"></div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import * as THREE from 'three';
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls';
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader';

const container = ref(null);
let scene, camera, renderer, controls, book, pages;

const createScene = () => {
  scene = new THREE.Scene();
  camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
  renderer = new THREE.WebGLRenderer({ antialias: true });
  renderer.setSize(window.innerWidth, window.innerHeight);
  renderer.setClearColor(0xf0f0f0);
  container.value.appendChild(renderer.domElement);

  camera.position.set(0, 0, 5);
  controls = new OrbitControls(camera, renderer.domElement);
  controls.enableDamping = true;
  controls.dampingFactor = 0.05;

  const ambientLight = new THREE.AmbientLight(0xffffff, 0.5);
  scene.add(ambientLight);

  const directionalLight = new THREE.DirectionalLight(0xffffff, 0.8);
  directionalLight.position.set(5, 5, 5);
  scene.add(directionalLight);
};

const loadBook = () => {
  const loader = new GLTFLoader();
  loader.load(
      '/assets/3d/book_model.glb',
      (gltf) => {
        book = gltf.scene;
        book.scale.set(2, 2, 2);
        scene.add(book);

        // Find and store references to book pages
        pages = [];
        book.traverse((child) => {
          if (child.isMesh && child.name.startsWith('Page')) {
            pages.push(child);
          }
        });

        // Sort pages by their names
        pages.sort((a, b) => {
          const aNum = parseInt(a.name.replace('Page', ''));
          const bNum = parseInt(b.name.replace('Page', ''));
          return aNum - bNum;
        });

        // Set up page turning animation
        setupPageTurning();
      },
      undefined,
      (error) => {
        console.error('An error occurred while loading the book model:', error);
      }
  );
};

const setupPageTurning = () => {
  let currentPage = 0;
  const turnDuration = 1000; // ms

  const turnPage = (direction) => {
    if (direction === 'next' && currentPage < pages.length - 1) {
      const page = pages[currentPage];
      new THREE.TWEEN.Tween(page.rotation)
          .to({ y: Math.PI }, turnDuration)
          .easing(THREE.TWEEN.Easing.Quadratic.InOut)
          .start();
      currentPage++;
    } else if (direction === 'prev' && currentPage > 0) {
      currentPage--;
      const page = pages[currentPage];
      new THREE.TWEEN.Tween(page.rotation)
          .to({ y: 0 }, turnDuration)
          .easing(THREE.TWEEN.Easing.Quadratic.InOut)
          .start();
    }
  };

  window.addEventListener('keydown', (event) => {
    if (event.key === 'ArrowRight') {
      turnPage('next');
    } else if (event.key === 'ArrowLeft') {
      turnPage('prev');
    }
  });

  renderer.domElement.addEventListener('click', (event) => {
    const rect = renderer.domElement.getBoundingClientRect();
    const x = ((event.clientX - rect.left) / rect.width) * 2 - 1;
    if (x > 0) {
      turnPage('next');
    } else {
      turnPage('prev');
    }
  });
};

const animate = () => {
  requestAnimationFrame(animate);
  controls.update();
  THREE.TWEEN.update();
  renderer.render(scene, camera);
};

const handleResize = () => {
  camera.aspect = window.innerWidth / window.innerHeight;
  camera.updateProjectionMatrix();
  renderer.setSize(window.innerWidth, window.innerHeight);
};

onMounted(() => {
  createScene();
  loadBook();
  animate();
  window.addEventListener('resize', handleResize);
});

onUnmounted(() => {
  window.removeEventListener('resize', handleResize);
  renderer.dispose();
});
</script>

<style scoped>
canvas {
  display: block;
}
</style>