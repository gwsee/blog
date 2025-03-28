<template>
  <div class="min-h-screen bg-transparent">
    <!-- 搜索区域 -->
    <div class="container mx-auto px-4 py-6">
      <div class="bg-[#f5f0e1] rounded-lg shadow-sm p-2 flex flex-wrap items-center">
        <div class="w-full md:w-1/5 p-1">
          <div class="relative">
            <!-- 关键字选择区域 - 类似 antd a-select -->
            <div
                class="w-full px-4 py-2 border border-[#e0d9c5] rounded-lg focus-within:ring-2 focus-within:ring-[#f8a13f] bg-[#faf7f0] h-[40px] cursor-text flex items-center gap-1 overflow-hidden"
                @click="focusKeywordInput"
            >
              <!-- 已选择的标签容器 - 限制为单行并添加省略号 -->
              <div class="flex items-center gap-1 overflow-hidden whitespace-nowrap flex-shrink-1 max-w-[calc(100%-60px)]">
                <!-- 显示可见标签和标签计数器 -->
                <div class="flex items-center gap-1 overflow-hidden">
                  <!-- 已选择的标签 -->
                  <div
                      v-for="(tag, index) in selectedTags.slice(0, visibleTagCount)"
                      :key="tag"
                      class="bg-[#7fbda0] text-white text-xs px-2 py-1 rounded-full flex items-center gap-1 flex-shrink-0"
                  >
                    {{ tag }}
                    <button
                        @click.stop="removeTag(tag)"
                        class="hover:text-[#e0d9c5] transition-colors"
                    >
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
                      </svg>
                    </button>
                  </div>

                  <!-- 标签数量指示器 - 确保它在标签旁边 -->
                  <div v-if="selectedTags.length > visibleTagCount" class="text-xs text-[#7fbda0] bg-white/20 px-2 py-1 rounded-full flex-shrink-0">
                    +{{ selectedTags.length - visibleTagCount }}
                  </div>
                </div>
              </div>

              <!-- 输入框 -->
              <input
                  ref="keywordInput"
                  type="text"
                  v-model="inputKeyword"
                  :placeholder="selectedTags.length === 0 ? '请输入关键字' : ''"
                  class="flex-1 bg-transparent border-none outline-none p-0 min-w-[50px] text-sm"
                  @keydown.enter="addKeyword"
                  @blur="addKeyword"
              />

              <!-- 清空按钮 -->
              <button
                  v-if="selectedTags.length > 0"
                  @click.stop="clearTags"
                  class="text-[#7fbda0] hover:text-[#6ca98c] ml-auto flex-shrink-0"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
          </div>
        </div>
        <div class="w-full md:w-3/5 p-1">
          <input
              type="text"
              v-model="searchContent"
              placeholder="请输入需要查询的内容"
              class="w-full px-4 py-2 border border-[#e0d9c5] rounded-lg focus:outline-none focus:ring-2 focus:ring-[#f8a13f] bg-[#faf7f0]"
          />
        </div>
        <div class="w-full md:w-1/5 p-1 flex space-x-2">
          <button
              @click="onSearch"
              class="flex-1 bg-[#7fbda0] hover:bg-[#6ca98c] text-white py-2 px-4 rounded-lg transition-colors"
          >
            查询
          </button>
          <button
              v-if="state.user?.id"
              @click="toRoute('/blog/manage/0')"
              class="flex-1 bg-[#f8a13f] hover:bg-[#e6912d] text-white py-2 px-4 rounded-lg transition-colors"
          >
            新增
          </button>
        </div>
      </div>
    </div>
    <!-- 三栏布局 -->
    <div class="container mx-auto px-4 pb-8 relative">
      <!-- 标签列表折叠按钮 - PC端左侧，移动端左下角 -->
      <button
          v-if="!showTagList"
          @click="toggleTagList"
          class="fixed z-10 md:fixed md:left-6 md:top-1/2 md:-translate-y-1/2 left-4 bottom-4 w-12 h-12 rounded-full bg-[#7fbda0] text-white shadow-lg flex items-center justify-center hover:bg-[#6ca98c] transition-colors"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h.01M7 12h.01M7 17h.01M11 7h6M11 12h6M11 17h6" />
        </svg>
      </button>

      <!-- 热门文章折叠按钮 - PC端右侧，移动端右下角 -->
      <button
          v-if="!showHotArticles"
          @click="toggleHotArticles"
          class="fixed z-10 md:fixed md:right-6 md:top-1/2 md:-translate-y-1/2 right-4 bottom-4 w-12 h-12 rounded-full bg-[#8a8f99] text-white shadow-lg flex items-center justify-center hover:bg-[#767b85] transition-colors"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
        </svg>
      </button>

      <div class="flex flex-wrap">
        <!-- 左侧标签列表 - 可折叠 -->
        <transition name="slide-fade">
          <div v-if="showTagList" class="w-full md:w-1/5 px-1 mb-6 md:mb-0">
            <div class="bg-[#7fbda0] bg-opacity-90 rounded-lg shadow-sm p-4 text-white">
              <div class="flex justify-between items-center mb-4">
                <h3 class="text-lg font-bold text-white">标签列表</h3>
                <button @click="toggleTagList" class="text-white hover:text-[#e0d9c5] transition-colors">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
                  </svg>
                </button>
              </div>
              <div class="space-y-2">
                <!-- 服务器标签 -->
                <div
                    v-for="tag in combinedTagList"
                    :key="tag.name"
                    class="flex justify-between items-center p-2 rounded cursor-pointer transition-colors"
                    :class="[
                      selectedTags.includes(tag.name) ? 'bg-[#6ca98c]' : 'hover:bg-[#6ca98c]',
                      tag.isCustom ? 'border border-dashed border-white/50' : ''
                    ]"
                    @click="selectTag(tag.name)"
                >
                  <span class="text-white" :class="{'font-medium': selectedTags.includes(tag.name)}">
                    {{ tag.name }}
                    <span v-if="tag.isCustom" class="text-xs ml-1">(自定义)</span>
                  </span>
                  <div class="flex items-center gap-2">
                    <!-- 选中标记 -->
                    <svg
                        v-if="selectedTags.includes(tag.name)"
                        xmlns="http://www.w3.org/2000/svg"
                        class="h-4 w-4 text-white"
                        viewBox="0 0 20 20"
                        fill="currentColor"
                    >
                      <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                    </svg>
                    <span class="bg-[#f5f0e1] text-[#7fbda0] text-xs px-2 py-1 rounded-full">
                      {{ tag.count }}
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </transition>

        <!-- 中间博客列表 -->
        <div
            class="px-1 mb-6 md:mb-0"
            :class="{
        'w-full': !showTagList && !showHotArticles,
        'w-full md:w-3/5': (showTagList && showHotArticles),
        'w-full md:w-4/5': (!showTagList || !showHotArticles) && !(!showTagList && !showHotArticles),
      }"
        >
          <!-- 博客卡片列表 -->
          <template v-if="blogList.length > 0">
            <div
                v-for="(item, index) in blogList"
                :key="index"
                class="bg-[#f5f0e1] rounded-lg shadow-sm mb-2 overflow-hidden border border-[#e0d9c5]"
            >
              <div class="flex p-4">
                <!-- 左侧头像 -->
                <div class="mr-4">
                  <img
                      :src="(item.cover && $fileFull(item.cover)) || defaultCover"
                      alt="Blog cover"
                      class="w-16 h-16 rounded-full object-cover border-2 border-[#7fbda0]"
                  />
                </div>

                <!-- 右侧内容 -->
                <div class="flex-1">
                  <!-- 标题区域 -->
                  <div class="flex justify-between items-center mb-2">
                    <h3
                        @click="toRoute('/blog/detail/'+item.id)"
                        class="text-lg font-medium truncate w-3/4 cursor-pointer hover:text-[#7fbda0] text-[#5a5a5a]"
                    >
                      {{ item.title }}
                    </h3>

                    <!-- 编辑/删除按钮 -->
                    <div v-if="state.user?.id === item.accountId" class="flex space-x-2">
                      <button
                          @click="toRoute('/blog/manage/'+item.id)"
                          class="text-xs text-[#7fbda0] hover:text-[#6ca98c] px-2 py-1 bg-[#e8e3d3] rounded-md"
                      >
                        编辑
                      </button>
                      <button @click="handleDel(item.id)" class="text-xs text-[#e05d5d] hover:text-[#c04545] px-2 py-1 bg-[#e8e3d3] rounded-md">
                        删除
                      </button>
                    </div>
                  </div>

                  <!-- 描述区域 - 只显示2行 -->
                  <div
                      @click="toRoute('/blog/detail/'+item.id)"
                      class="h-12 overflow-hidden text-[#6b6b6b] text-sm mb-2 cursor-pointer line-clamp-2"
                  >
                    {{ item.description || '这个人很懒什么都没写' }}
                  </div>

                  <!-- 标签和操作区域 -->
                  <div class="flex justify-between items-center">
                    <!-- 标签 -->
                    <div class="w-1/2 flex flex-wrap gap-1">
                  <span
                      v-for="tag in item.tags"
                      :key="tag"
                      class="px-2 py-0.5 text-xs bg-[#7fbda0] text-white rounded-full"
                  >
                    {{ tag }}
                  </span>
                    </div>

                    <!-- 点赞/收藏按钮 -->
                    <div class="flex space-x-4">
                      <button class="flex items-center text-[#f8a13f] hover:text-[#e6912d]">
                        <HeartOutlined class="mr-1" v-if="false" />
                        <span class="text-xs">{{ item.browseNum || '--' }}</span>
                      </button>
                      <button class="text-[#7fbda0] hover:text-[#6ca98c]">
                        <LikeOutlined  v-if="false"  />
                      </button>
                      <button class="text-[#e05d5d] hover:text-[#c04545]">
                        <DislikeOutlined  v-if="false"  />
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- 分页 - 有数据时显示 -->
            <div class="flex justify-center py-4 mt-4">
              <div class="flex items-center space-x-2">
                <button
                    @click="current > 1 ? current-- : null"
                    :disabled="current === 1"
                    class="px-3 py-1 rounded border border-[#e0d9c5] bg-[#f5f0e1] text-[#7fbda0] disabled:opacity-50"
                >
                  上一页
                </button>

                <span class="text-[#5a5a5a]">{{ current }} / {{ Math.ceil(total / pageSize) || 1 }}</span>

                <button
                    @click="current < Math.ceil(total / pageSize) ? current++ : null"
                    :disabled="current >= Math.ceil(total / pageSize) || total === 0"
                    class="px-3 py-1 rounded border border-[#e0d9c5] bg-[#f5f0e1] text-[#7fbda0] disabled:opacity-50"
                >
                  下一页
                </button>

                <select v-model="pageSize" class="px-2 py-1 border border-[#e0d9c5] bg-[#f5f0e1] rounded text-[#5a5a5a]">
                  <option v-for="size in pageSizeOptions" :key="size" :value="Number(size)">
                    {{ size }} 条/页
                  </option>
                </select>
              </div>
            </div>
          </template>

          <!-- 暂无数据提示 -->
          <div v-else class="bg-[#f5f0e1] rounded-lg shadow-sm p-8 text-center">
            <div class="flex flex-col items-center justify-center py-12">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 text-[#7fbda0] mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
              </svg>
              <h3 class="text-xl font-bold text-[#5a5a5a] mb-2">暂无数据</h3>
              <p class="text-[#6b6b6b] mb-6">没有找到符合条件的博客文章</p>
              <button
                  @click="resetSearch"
                  class="px-4 py-2 bg-[#7fbda0] hover:bg-[#6ca98c] text-white rounded-lg transition-colors"
              >
                重置搜索
              </button>
            </div>
          </div>
        </div>

        <!-- 右侧热门文章 - 可折叠 -->
        <transition name="slide-fade-right">
          <div v-if="showHotArticles" class="w-full md:w-1/5 px-1">
            <div class="bg-[#8a8f99] bg-opacity-90 rounded-lg shadow-sm p-4 text-white">
              <div class="flex justify-between items-center mb-4">
                <h3 class="text-lg font-bold text-white">热门文章</h3>
                <button @click="toggleHotArticles" class="text-white hover:text-[#e0d9c5] transition-colors">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
                  </svg>
                </button>
              </div>
              <div class="h-[29.5rem] overflow-hidden relative" ref="scrollContainer">
                <div class="space-y-4 animate-scroll">
                  <div
                      v-for="(article, index) in hotArticles"
                      :key="index"
                      class="p-2 hover:bg-[#767b85] rounded cursor-pointer transition-colors"
                      @click="toRoute('/blog/detail/'+article.id)"
                  >
                    <div class="text-sm font-medium text-white mb-1 truncate">{{ article.title }}</div>
                    <div class="text-xs text-[#e0e0e0]">{{ $formatDate(article.createdAt) }}</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </transition>
      </div>
    </div>
  </div>
</template>

<script setup>
import {ref, watch, onMounted, computed, nextTick} from 'vue';
import {useRoute, useRouter} from 'vue-router';
import {blogList as fetchBlogList, blogTags, blogHot,blogDel} from '@/api/blog';
import {useAuthStore} from '@/store/auth';
import defaultCover from '@/assets/image/default-cover.jpg';
import { HeartOutlined, LikeOutlined, DislikeOutlined } from '@ant-design/icons-vue';
// 路由
const router = useRouter();
const toRoute = (path) => router.push(path);

// 状态管理
const {state} = useAuthStore();

// 分页
const current = ref(1);
const pageSize = ref(5);
const total = ref(0);
const pageSizeOptions = ref(['5', '10', '20', '30', '40', '50']);

// 搜索
const inputKeyword = ref(''); // 输入框中的关键字
const searchContent = ref('');
const selectedTags = ref([]);
const keywordInput = ref(null);

// 计算可见标签数量
const visibleTagCount = computed(() => {
  // 根据容器宽度动态计算可见标签数量
  // 这里设置一个默认值，实际使用时可能需要根据容器宽度动态调整
  return 2; // 只显示前2个标签，可以根据需要调整
});

// 数据
const blogList = ref([]);
const loading = ref(false);

// 侧边栏显示状态
const showTagList = ref(true);
const showHotArticles = ref(true);

// 标签列表 - 分为服务器标签和自定义标签
const serverTagList = ref([]); // 服务器返回的标签列表
const customTagList = ref([]); // 用户自定义的标签列表

// 合并标签列表 - 服务器标签 + 自定义标签
const combinedTagList = computed(() => {
  // 先获取所有服务器标签的名称
  const serverTagNames = serverTagList.value.map(tag => tag.name);

  // 过滤出不在服务器标签中的自定义标签
  const filteredCustomTags = customTagList.value.filter(tag =>
      !serverTagNames.includes(tag.name) && selectedTags.value.includes(tag.name)
  );

  // 合并两个数组
  return [...serverTagList.value, ...filteredCustomTags];
});

// 热门文章
const hotArticles = ref([]);

// 添加关键字
const addKeyword = () => {
  if (inputKeyword.value.trim() !== '') {
    const newTag = inputKeyword.value.trim();

    // 避免重复添加到已选标签
    if (!selectedTags.value.includes(newTag)) {
      selectedTags.value.push(newTag);

      // 检查是否存在于服务器标签列表中
      const existsInServerTagList = serverTagList.value.some(tag => tag.name === newTag);

      // 检查是否存在于自定义标签列表中
      const existsInCustomTagList = customTagList.value.some(tag => tag.name === newTag);

      // 如果不存在于任何标签列表中，则添加到自定义标签列表
      if (!existsInServerTagList && !existsInCustomTagList) {
        customTagList.value.push({
          name: newTag,
          count: 0,
          isCustom: true // 标记为自定义标签
        });
      }
    }

    inputKeyword.value = '';
  }
};

// 聚焦关键字输入框
const focusKeywordInput = () => {
  nextTick(() => {
    keywordInput.value.focus();
  });
};

// 移除单个标签
const removeTag = (tag) => {
  const index = selectedTags.value.indexOf(tag);
  if (index !== -1) {
    selectedTags.value.splice(index, 1);
  }
  // clearCustomTag(tag);
};
// const clearCustomTag = (tag) => {
//   const index = customTagList.value.findIndex(tag => tag.name === tag);
//   if (index !== -1) {
//     customTagList.value.splice(index, 1);
//   }
//   console.log(customTagList);
// }
// 清空所有标签
const clearTags = () => {
  selectedTags.value = [];
  customTagList.value = [];
};
const handleDel=(id)=>{
  blogDel({id:id}).then(res=>{

  }).finally(()=>{
    fetchData()
  })
}
// 切换标签列表显示
const toggleTagList = () => {
  showTagList.value = !showTagList.value;
  let id = state.user?.id || '0';
  // 保存状态到本地存储
  localStorage.setItem('blogTagListVisible' + id, showTagList.value);
};

// 切换热门文章显示
const toggleHotArticles = () => {
  showHotArticles.value = !showHotArticles.value;
  let id = state.user?.id || '0';
  // 保存状态到本地存储
  localStorage.setItem('blogHotArticlesVisible' + id, showHotArticles.value);
};

// 选择标签
const selectTag = (tag) => {
  const index = selectedTags.value.indexOf(tag);
  if (index === -1) {
    selectedTags.value.push(tag);
  } else {
    selectedTags.value.splice(index, 1);
  }
};

// 搜索
const onSearch = () => {
  // 如果输入框有内容，先添加为标签
  if (inputKeyword.value.trim() !== '') {
    addKeyword();
  }

  if (current.value === 1) {
    fetchData();
  } else {
    current.value = 1;
  }
};

// 重置搜索
const resetSearch = () => {
  inputKeyword.value = '';
  searchContent.value = '';
  selectedTags.value = [];
  customTagList.value = [];
  current.value = 1;
  fetchData();
};

// 获取标签
const fetchTags = (params) => {
  blogTags(Object.assign({}, params, {page: {pageSize: 10}})).then(res => {
    if (res && res.code === 200) {
      const tags = res.data && res.data.tags || [];
      const data = [];
      tags.forEach(tag => {
        data.push({name: tag.name, count: tag.num, isCustom: false});
      });
      serverTagList.value = data;
    }
  });
};

const fetchHot = () => {
  blogHot({pageSize: 10}).then(res => {
    if (res && res.code === 200) {
      hotArticles.value = res.data.list || [];
    }
  });
};

// 获取博客列表
const fetchData = () => {
  const params = {
    page: {
      pageNum: current.value,
      pageSize: pageSize.value
    },
    title: searchContent.value,
    keyword: inputKeyword.value,
    tags: selectedTags.value
  };

  if (params.page.pageNum === 1) {
    fetchTags(params);
  }

  loading.value = true;
  fetchBlogList(params)
      .then(res => {
        if (res && res.code === 200) {
          blogList.value = res.data.list || [];
          total.value = res.data.total || 0;
        }
      })
      .finally(() => {
        loading.value = false;
      });
};

// 监听分页变化
watch(pageSize, () => {
  if (current.value === 1) {
    fetchData();
  } else {
    current.value = 1;
  }
});

watch(current, () => {
  fetchData();
});

// 监听selectedTags变化，自动触发查询
watch(selectedTags, () => {
  if (current.value === 1) {
    fetchData();
  } else {
    current.value = 1;
  }
}, {deep: true});
const getRouter = ()=>{
  const route = useRoute();
  let q = route.query&&route.query.q;
  if(q){
    searchContent.value = q
  }
}
// 初始化
onMounted(() => {
  getRouter()
  let id = state.user?.id || '0';
  // 从本地存储恢复侧边栏状态
  const savedTagListVisible = localStorage.getItem('blogTagListVisible' + id);
  const savedHotArticlesVisible = localStorage.getItem('blogHotArticlesVisible' + id);

  if (savedTagListVisible !== null) {
    showTagList.value = savedTagListVisible === 'true';
  }

  if (savedHotArticlesVisible !== null) {
    showHotArticles.value = savedHotArticlesVisible === 'true';
  }

  // 从本地存储恢复自定义标签
  const savedCustomTags = localStorage.getItem('blogCustomTags' + id);
  if (savedCustomTags !== null) {
    try {
      customTagList.value = JSON.parse(savedCustomTags);
    } catch (e) {
      console.error('解析自定义标签失败', e);
      customTagList.value = [];
    }
  }

  fetchHot();
  fetchData();
});

// 监听自定义标签变化，保存到本地存储
watch(customTagList, () => {
  let id = state.user?.id || '0';
  localStorage.setItem('blogCustomTags' + id, JSON.stringify(customTagList.value));
}, { deep: true });
</script>

<style>
@tailwind base;
@tailwind components;
@tailwind utilities;

body {
  background-size: cover;
  background-attachment: fixed;
  background-position: center;
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.animate-scroll {
  animation: scroll 20s linear infinite;
}

@keyframes scroll {
  0% {
    transform: translateY(0);
  }
  100% {
    transform: translateY(-50%);
  }
}

/* 当鼠标悬停时暂停滚动 */
.animate-scroll:hover {
  animation-play-state: paused;
}

/* 左侧侧边栏过渡动画 */
.slide-fade-enter-active,
.slide-fade-leave-active {
  transition: all 0.3s ease;
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(-20px);
  opacity: 0;
}

/* 右侧侧边栏过渡动画 */
.slide-fade-right-enter-active,
.slide-fade-right-leave-active {
  transition: all 0.3s ease;
}

.slide-fade-right-enter-from,
.slide-fade-right-leave-to {
  transform: translateX(20px);
  opacity: 0;
}
</style>

