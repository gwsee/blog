<template>
  <div class="relative w-full h-screen">
    <canvas ref="canvas" class="w-full h-full"></canvas>

    <!-- Loading Screen -->
    <div v-if="!sceneReady" class="absolute inset-0 flex items-center justify-center bg-orange-500">
      <div class="text-white text-2xl">Loading... {{ Math.round(loadingProgress) }}%</div>
    </div>

    <!-- Navigation UI -->
    <div v-show="sceneReady" class="absolute bottom-4 left-4 space-y-2">
      <button
          v-for="section in sections"
          :key="section.id"
          @click="navigateToSection(section.id)"
          class="px-4 py-2 bg-white/90 rounded-full text-orange-500 hover:bg-white transition-colors"
      >
        {{ section.title }}
      </button>
    </div>

    <!-- Info Panel -->
    <div
        v-if="activeSection && sceneReady"
        class="absolute top-4 right-4 w-80 bg-white/90 p-6 rounded-lg shadow-lg"
    >
      <h2 class="text-xl font-bold text-orange-500 mb-4">{{ activeSection.title }}</h2>
      <div class="space-y-4">
        <div v-for="(item, index) in activeSection.content" :key="index">
          <h3 class="font-semibold">{{ item.title }}</h3>
          <p class="text-sm text-gray-600">{{ item.description }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue'
import * as THREE from 'three'
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader'
import { OrbitControls } from 'three/examples/jsm/controls/OrbitControls'

// Scene setup
const canvas = ref(null)
const sceneReady = ref(false)
const loadingProgress = ref(0)
const activeSection = ref(null)

// Portfolio content
const sections = ref([
  {
    id: 'about',
    title: 'About Me',
    position: new THREE.Vector3(-10, 0, 0),
    content: [
      {
        title: 'Introduction',
        description: 'A passionate developer with expertise in Vue.js and 3D web experiences.'
      }
    ]
  },
  {
    id: 'experience',
    title: 'Experience',
    position: new THREE.Vector3(0, 0, 0),
    content: [
      {
        title: 'Senior Frontend Developer',
        description: '2020-Present: Leading development of interactive web applications'
      },
      {
        title: 'Web Developer',
        description: '2018-2020: Full-stack development with Vue.js and Node.js'
      }
    ]
  },
  {
    id: 'projects',
    title: 'Projects',
    position: new THREE.Vector3(10, 0, 0),
    content: [
      {
        title: '3D Portfolio',
        description: 'Interactive portfolio website built with Three.js and Vue'
      },
      {
        title: 'E-commerce Platform',
        description: 'Modern shopping experience with Vue 3 and Tailwind CSS'
      }
    ]
  }
])

// Three.js variables
let scene, camera, renderer, controls
let mixer, clock
const models = new Map()

// Initialize Three.js scene
const initScene = () => {
  // Scene
  scene = new THREE.Scene()
  scene.background = new THREE.Color('#ff7b4d') // Orange background

  // Camera
  camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000)
  camera.position.set(0, 5, 20)

  // Renderer
  renderer = new THREE.WebGLRenderer({
    canvas: canvas.value,
    antialias: true
  })
  renderer.setSize(window.innerWidth, window.innerHeight)
  renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2))

  // Controls
  controls = new OrbitControls(camera, renderer.domElement)
  controls.enableDamping = true
  controls.dampingFactor = 0.05
  controls.maxPolarAngle = Math.PI / 2

  // Lighting
  const ambientLight = new THREE.AmbientLight(0xffffff, 0.5)
  scene.add(ambientLight)

  const directionalLight = new THREE.DirectionalLight(0xffffff, 1)
  directionalLight.position.set(5, 5, 5)
  scene.add(directionalLight)

  // Clock for animations
  clock = new THREE.Clock()

  // Load 3D models
  loadModels()
}

// Load 3D models and setup scene
const loadModels = () => {
  const loader = new GLTFLoader()
  const totalModels = sections.value.length
  let loadedModels = 0

  sections.value.forEach(section => {
    // Create platform for each section
    const platform = new THREE.Mesh(
        new THREE.BoxGeometry(5, 0.5, 5),
        new THREE.MeshStandardMaterial({ color: 0xffffff })
    )
    platform.position.copy(section.position)
    scene.add(platform)

    // Add simple figure placeholder (can be replaced with actual GLTF models)
    const figure = new THREE.Mesh(
        new THREE.CylinderGeometry(0.5, 0.5, 2, 8),
        new THREE.MeshStandardMaterial({ color: 0xffa07a })
    )
    figure.position.copy(section.position)
    figure.position.y += 1.25
    scene.add(figure)

    models.set(section.id, { platform, figure })

    loadedModels++
    loadingProgress.value = (loadedModels / totalModels) * 100
  })

  sceneReady.value = true
}

// Animation loop
const animate = () => {
  if (!sceneReady.value) return

  const delta = clock.getDelta()
  if (mixer) mixer.update(delta)

  controls.update()
  renderer.render(scene, camera)
  requestAnimationFrame(animate)
}

// Navigation
const navigateToSection = (sectionId) => {
  const section = sections.value.find(s => s.id === sectionId)
  if (!section) return

  activeSection.value = section

  // Animate camera to new position
  const targetPosition = section.position.clone()
  targetPosition.add(new THREE.Vector3(0, 5, 10))

  gsap.to(camera.position, {
    duration: 1.5,
    x: targetPosition.x,
    y: targetPosition.y,
    z: targetPosition.z,
    ease: 'power2.inOut'
  })
}

// Resize handler
const handleResize = () => {
  if (!camera || !renderer) return

  camera.aspect = window.innerWidth / window.innerHeight
  camera.updateProjectionMatrix()
  renderer.setSize(window.innerWidth, window.innerHeight)
}

// Lifecycle hooks
onMounted(() => {
  initScene()
  animate()
  window.addEventListener('resize', handleResize)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize)
  if (renderer) renderer.dispose()
})
</script>