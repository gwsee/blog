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
                    @click="toRoute('/travel/info/'+travel.id)"
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

                    <!-- Title and Time Header -->
                    <div class="absolute top-0 left-0 right-0 p-4 flex justify-between items-center bg-gradient-to-b from-black/70 to-transparent">
                      <h3 class="text-lg font-medium text-white truncate flex-1 mr-2">
                        {{ travel.title }}
                      </h3>
                      <span class="text-xs text-white/80 whitespace-nowrap">
                        {{ $formatDate(travel.createdAt || travel.updatedAt) }}
                      </span>
                    </div>

                    <!-- travel Description (visible on hover) -->
                    <div class="absolute bottom-0 left-0 right-0 p-4 bg-gradient-to-t from-black/70 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-200">
                      <p class="text-sm text-white/90 line-clamp-3">{{ travel.description }}</p>
                    </div>

                    <!-- Action Buttons (always visible) -->
                    <div class="absolute  inset-0 flex items-center  space-x-4 justify-center opacity-0 group-hover:opacity-100 transition-opacity duration-200">
                      <a-button
                          class="action-button hover:scale-110 transition-transform"
                          :class="[
                            key % 2 !== 0 ? 'bg-white hover:bg-white/90' : 'bg-black hover:bg-black/90',
                            travel.isCollect ? '!border-2 !border-red-500' : ''
                          ]"
                          shape="circle"
                          size="large"
                          @click.stop="travelCollectAc(travel)"
                      >
                        <template #icon>
                          <StarFilled v-if="travel.isCollect"    :class="[
                              key % 2 !== 0 ? 'text-black' : 'text-white',
                              travel.isCollect ? '!text-red-500' : ''
                            ]"/>
                          <StarOutlined v-else   :class="[
                              key % 2 !== 0 ? 'text-black' : 'text-white',
                              travel.isCollect ? '!text-red-500' : ''
                            ]"/>
                        </template>
                      </a-button>
                      <a-button
                          class="action-button hover:scale-110 transition-transform"
                          :class="[
                            key % 2 !== 0 ? 'bg-white hover:bg-white/90' : 'bg-black hover:bg-black/90',
                            travel.isThumb ? '!border-2 !border-red-500' : ''
                          ]"
                          shape="circle"
                          size="large"
                          @click.stop="travelThumbAc(travel)"
                      >
                        <template #icon>
                          <HeartFilled v-if="travel.isThumb "  :class="[
                              key % 2 !== 0 ? 'text-black' : 'text-white',
                              travel.isThumb ? '!text-red-500' : ''
                            ]" />
                          <HeartOutlined v-else  :class="[
                              key % 2 !== 0 ? 'text-black' : 'text-white',
                              travel.isThumb ? '!text-red-500' : ''
                            ]"/>
                        </template>
                      </a-button>
                      <a-button
                          v-if="state.user && state.user.id === travel.account.id"
                          class="action-button hover:scale-110 transition-transform"
                          :class="key % 2 !== 0 ? 'bg-white hover:bg-white/90' : 'bg-black hover:bg-black/90'"
                          shape="circle"
                          size="large"
                          @click.stop="editTravel(travel)"
                      >
                        <template #icon>
                          <EditOutlined :class="key % 2 !== 0 ? 'text-black' : 'text-white'" />
                        </template>
                      </a-button>
                      <a-button
                          v-if="state.user && state.user.id === travel.account.id"
                          class="action-button hover:scale-110 transition-transform"
                          :class="key % 2 !== 0 ? 'bg-white hover:bg-white/90' : 'bg-black hover:bg-black/90'"
                          shape="circle"
                          size="large"
                          @click.stop="editTravel(travel)"
                      >
                        <template #icon>
                          <DeleteOutlined :class="key % 2 !== 0 ? 'text-black' : 'text-white'" />
                        </template>
                      </a-button>
                    </div>
                  </div>

                  <!-- User Info Footer -->
                  <div :class="[
                    'p-4 flex items-center justify-between',
                    key % 2 !== 0 ? 'bg-white' : 'bg-black'
                  ]">
                    <div class="flex items-center gap-2">
                      <a-avatar :src="travel.account.avatar||avatar" :size="24" />
                      <span class="text-sm font-medium truncate max-w-[20ch]" :class="key % 2 !== 0 ? 'text-black' : 'text-white'">
                        {{ travel.account.nickname }}
                      </span>
                      <a-tag
                          v-if="state.user&&state.user.id === travel.account.id"
                          class="border-0 text-xs"
                          :class="key % 2 !== 0 ? 'bg-black text-white' : 'bg-white text-black'"
                      >
                        ME
                      </a-tag>
                    </div>
                    <div class="flex items-center gap-3">
                       <span class="flex items-center gap-1"  :class="[
                            key % 2 !== 0 ? 'text-black' : 'text-white',
                            travel.isCollect ? 'text-red-500' : '',
                          ]" >
                          <StarFilled v-if="travel.isCollect "  />
                          <StarOutlined v-else  />
                        {{ travel.collectNum>999?'999+':(travel.collectNum||0) }}
                      </span>
                      <span class="flex items-center gap-1" :class="[
                            key % 2 !== 0 ? 'text-black' : 'text-white',
                            travel.isThumb ? 'text-red-500' : '',
                          ]" >
                         <HeartFilled v-if="travel.isThumb "   />
                          <HeartOutlined v-else   />
                        {{ travel.thumbNum>999?'999+':(travel.thumbNum||0) }}
                      </span>
                      <span class="flex items-center gap-1" :class="key % 2 !== 0 ? 'text-black' : 'text-white'">
                        <EyeOutlined />
                        {{ travel.browseNum>999?'999+':(travel.browseNum||0) }}
                      </span>
                    </div>
                  </div>
                </div>
              </div>
            </template>
          </div>
          <div v-if="false&&total>pageSize">
            <!--后续做成滚动刷新 -->
            <a-pagination style="float: right" v-model:current="current" :total="total" :page-size-options="pageSizeOptions"  v-model:pageSize="pageSize" />
          </div>
          <div
              v-if="state.user&&state.user.id"
              class="sticky float-right bottom-2 right-2 w-12 h-12 text-white flex items-center justify-center rounded-full right">
            <a-button
                class="action-button !bg-white/90 hover:!bg-white"
                shape="circle"
                size="large"
                style="background-color: gray"
                @click="toRoute('/travel/manage/0')"
            >
              <template #icon>
                <PlusOutlined :class=" 'text-black'" />
              </template>
            </a-button>
          </div>

          <!-- Pagination and New Button Container -->
          <div class="mt-8 flex justify-between items-center" v-if="false">
            <a-pagination
                v-show="total>0"
                v-model:current="current" :total="total" :page-size-options="pageSizeOptions"  v-model:pageSize="pageSize"
            />
            <a-button
                class="action-button !bg-white/90 hover:!bg-white"
                shape="circle"
                size="large"
                style="background-color: gray"
                @click="toRoute('/travel/manage/0')"
            >
              <template #icon>
                <PlusOutlined :class=" 'text-black'" />
              </template>
            </a-button>
          </div>
        </div>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script setup>
import {ref, computed, onMounted, watch} from 'vue'
import {travelList, travelDel, travelCollect, travelThumb} from "@/api/travel";
import {filePrefix} from "@/api/tool";
import { useAuthStore  } from '@/store/auth'
const { state } = useAuthStore();
import {
  HeartOutlined,
  HeartFilled,
  StarOutlined,
  StarFilled,
  EyeOutlined,
  SaveOutlined,
  SaveFilled,
  EditOutlined,
  PlusOutlined,
  DeleteOutlined
} from '@ant-design/icons-vue'
import avatar from "@/assets/image/default-avatar.png"
import { useRouter } from "vue-router";
import { Empty } from 'ant-design-vue';

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
const pageSizeOptions = ref([ '1','4', '20', '40', '100']);
const current = ref(1)
const total = ref(0)
const pageSize = ref(1)
const loading = ref(false)
const params = ref({my:false,title:undefined,myCollect:false,sort:"",myThumb:false,description:""})
const travels = ref([])
watch(pageSize, () => {
  if(current.value===1){
    list()
  }else{
    current.value = 1
  }
});
watch(current, () => {
  list()
});
const list =()=>{
  travelList(Object.assign({},params.value,{"page":{"pageNum":current.value,"pageSize":pageSize.value}})).then(res=>{
    if(res&&res.code===200){
      travels.value = res.data.list||[]
      total.value = res.data.total||0
    }
  })
}


const travelCollectAc = async (travel) => {
  try {
    travel.collectNum=travel.collectNum||0
    travel.isCollect = !travel.isCollect
    travel.collectNum +=  travel.isCollect? 1 : -1
    await travelCollect({id:travel.id,do:travel.isCollect})
  } catch (error) {
    console.error('Error toggling favorite:', error)
  }
}

const travelThumbAc = async (travel) => {
  try {
    travel.thumbNum=  travel.thumbNum||0
    travel.isThumb = !travel.isThumb
    travel.thumbNum += travel.isThumb ? 1 : -1
    await travelThumb({id:travel.id,do:travel.isThumb})

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

