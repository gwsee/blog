<template>
  <div class="min-h-screen">
    <a-layout style="background-color: transparent">
      <a-layout-content class="p-4 md:p-8">
        <div class="relative max-w-7xl mx-auto">
          <!-- Grid Container -->
          <div :class="gridContainerClasses">
            <template v-if="travels.length === 0">
              <div class="text-center">
                <a-empty
                    description="暂无数据"
                    :image="Empty.PRESENTED_IMAGE_SIMPLE"
                >
                </a-empty>
              </div>
            </template>
            <template v-else>
              <div
                  v-for="(travel,key) in travels"
                  :key="travel.id"
                  :class="travelCardClasses"
              >
                <!-- travel Card -->
                <div
                    class="travel-card group relative overflow-hidden rounded-lg shadow-lg"
                    :class="{'bg-black': key % 2 !== 0}"
                >
                  <!-- travel Image Container with Fixed Aspect Ratio -->
                  <div class="relative w-full pb-[75%]">
                    <img
                        v-if="travel.photo"
                        :src="filePrefix+travel.photo"
                        :alt="travel.title"
                        class="absolute inset-0 w-full h-full object-cover"
                    />
                    <div
                        v-else
                        class="absolute inset-0 bg-gray-100 flex items-center justify-center"
                    >
                      <div class="text-4xl text-gray-400">No Photo</div>
                    </div>

                    <!-- Creation/Update Time -->
                    <div class="absolute top-2 right-2 bg-black/50 text-white/80 px-2 py-1 rounded text-xs">
                      {{ $formatDate(travel.createdAt || travel.updatedAt) }}
                    </div>

                    <!-- travel Title -->
                    <div class="absolute top-0 left-0 p-4">
                      <h3 class="text-lg font-medium text-white max-w-[12em] truncate">
                        {{ travel.title }}
                      </h3>
                    </div>

                    <!-- travel Description (visible on hover) -->
                    <div class="absolute bottom-0 left-0 right-0 p-4 bg-gradient-to-t from-black/70 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-200">
                      <p class="text-sm text-white/90 line-clamp-3">{{ travel.description }}</p>
                    </div>

                    <!-- Hover Actions -->
                    <div
                        class="absolute top-0 left-0 right-0 bottom-0 bg-black/40 opacity-0
                           group-hover:opacity-100 transition-opacity duration-200 flex
                           items-center justify-center gap-2"
                    >
                      <a-button
                          class="action-button hover:scale-110 transition-transform"
                          :class="[
                            key % 2 !== 0 ? '!bg-white/90 hover:!bg-white' : '!bg-black/90 hover:!bg-black',
                            travel.isFavorited ? 'border-2 !border-yellow-500' : ''
                          ]"
                          shape="circle"
                          size="large"
                          @click="toggleFavoriteAc(travel)"
                      >
                        <template #icon>
                          <component
                              :is="travel.isFavorited ? 'SaveFilled' : 'SaveOutlined'"
                              :class="[
                              {'text-black': key % 2 !== 0, 'text-white': key % 2 === 0},
                              travel.isFavorited ? 'text-yellow-500' : ''
                            ]"
                          />
                        </template>
                      </a-button>
                      <a-button
                          class="action-button hover:scale-110 transition-transform"
                          :class="[
                            key % 2 !== 0 ? '!bg-white/90 hover:!bg-white' : '!bg-black/90 hover:!bg-black',
                            travel.isLiked ? 'border-2 !border-red-500' : ''
                          ]"
                          shape="circle"
                          size="large"
                          @click="toggleLikeAc(travel)"
                      >
                        <template #icon>
                          <component
                              :is="travel.isLiked ? 'HeartFilled' : 'HeartOutlined'"
                              :class="[
                              {'text-black': key % 2 !== 0, 'text-white': key % 2 === 0},
                              travel.isLiked ? 'text-red-500' : ''
                            ]"
                          />
                        </template>
                      </a-button>
                      <a-button
                          v-if="state.user && state.user.id === travel.account.id"
                          class="action-button hover:scale-110 transition-transform"
                          :class="key % 2 !== 0 ? '!bg-white/90 hover:!bg-white' : '!bg-black/90 hover:!bg-black'"
                          shape="circle"
                          size="large"
                          @click="editTravel(travel)"
                      >
                        <template #icon>
                          <EditOutlined :class="{'text-black': key % 2 !== 0, 'text-white': key % 2 === 0}" />
                        </template>
                      </a-button>
                    </div>
                  </div>

                  <!-- travel Info -->
                  <div
                      class="absolute bottom-0 left-0 right-0 p-4 bg-gradient-to-t from-black/90 via-black/70 to-transparent"
                  >
                    <div class="flex items-center justify-between text-white">
                      <div class="flex items-center gap-2">
                        <a-avatar :src="travel.account.avatar||avatar" :size="24" />
                        <span class="text-sm font-medium">
                          {{ travel.account.nickname }}
                        </span>
                        <a-tag
                            v-if="state.user&&state.user.id === travel.account.id"
                            class="border-0 text-xs bg-white/20"
                        >
                          ME
                        </a-tag>
                      </div>
                      <div class="flex items-center gap-3">
                        <span class="flex items-center gap-1">
                          <component
                              :is="travel.isLiked ? 'HeartFilled' : 'HeartOutlined'"
                              :class="travel.isLiked ? 'text-red-500' : ''"
                          />
                          {{ travel.thumbNum>999?'999+':(travel.thumbNum) }}
                        </span>
                        <span class="flex items-center gap-1">
                          <EyeOutlined />
                          {{ travel.browseNum>999?'999+':(travel.browseNum) }}
                        </span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </template>
          </div>
          <div
              class="sticky float-right bottom-2 right-2 w-12 h-12 bg-blue-500 text-white flex items-center justify-center rounded-full right"
          >
            <a-button
                class="action-button !bg-white/90 hover:!bg-white"
                shape="circle"
                size="large"
                @click="toRoute('/travel/manage/0')"
            >
              NEW
            </a-button>
          </div>
        </div>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import {travelList, travelDel, toggleFavorite, toggleLike} from "@/api/travel";
import {filePrefix} from "@/api/tool";
import { useAuthStore  } from '@/store/auth'
import {
  HeartOutlined,
  HeartFilled,
  EyeOutlined,
  SaveOutlined,
  SaveFilled,
  EditOutlined
} from '@ant-design/icons-vue'
import avatar from "@/assets/image/default-avatar.png"
import { useRouter } from "vue-router";
import { Empty } from 'ant-design-vue';
const { state } = useAuthStore();
const router = useRouter();
const toRoute = (path) => {
  router.push(path)
}

const gridContainerClasses = computed(() => {
  const count = travels.value.length;
  if (count === 0) {
    return 'flex justify-center items-center min-h-[50vh]';
  } else if (count === 1) {
    return 'flex justify-center items-center min-h-[50vh]';
  } else if (count === 2) {
    return 'grid grid-cols-1 md:grid-cols-2 gap-4 justify-items-center';
  } else if (count <= 4) {
    return 'grid grid-cols-1 sm:grid-cols-2 gap-4';
  } else {
    return 'grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4';
  }
})
const travelCardClasses = computed(() => {
  const count = travels.value.length;
  if (count === 1) {
    return 'w-full max-w-2xl';
  } else if (count === 2) {
    return 'w-full max-w-xl';
  } else if (count <= 4) {
    return 'w-full max-w-lg';
  } else {
    return 'w-full';
  }
})
onMounted(()=>{
  list()
})
const pageSizeOptions = ref([ '20', '30', '40', '50']);
const current = ref(1)
const total = ref(0)
const pageSize = ref(20)
const loading = ref(false)
const params = ref({my:false,title:undefined,myCollect:false,sort:"",myThumb:false,description:""})
const travels = ref([])
const list =()=>{
  travelList(Object.assign({},params.value,{"page":{"pageNum":current.value,"pageSize":pageSize.value}})).then(res=>{
    if(res&&res.code===200){
      travels.value = res.data.list||[]
      total.value = res.data.total||0
    }
  })
}


const toggleFavoriteAc = async (travel) => {
  try {
    await toggleFavorite(travel.id)
    travel.isFavorited = !travel.isFavorited
  } catch (error) {
    console.error('Error toggling favorite:', error)
  }
}

const toggleLikeAc = async (travel) => {
  try {
    await toggleLike(travel.id)
    travel.isLiked = !travel.isLiked
    travel.thumbNum += travel.isLiked ? 1 : -1
  } catch (error) {
    console.error('Error toggling like:', error)
  }
}

const editTravel = (travel) => {
  router.push(`/travel/manage/${travel.id}`)
}
</script>

<style scoped>
.travel-card {
  transition: transform 0.2s ease;
}

.travel-card:hover {
  transform: translateY(-2px);
}

.action-button {
  transition: all 0.2s ease;
}

/* Override Ant Design button styles */
:deep(.ant-btn-circle) {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 44px;
  height: 44px;
}

.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>

