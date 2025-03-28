<script setup>
import { ref, computed, onMounted, onUnmounted, watch, reactive } from 'vue'
import {messages} from "@/api/index.js";
import {memoList,memoSave,memoDelete,memoDone,
  todoList,todoSave,todoDone,todoDelete,todoDoneDelete,
  todoDoneList,todoRecordDelete} from "@/api/palaces.js";
import {
  MinusCircleOutlined,
  PlusCircleOutlined,
  CheckCircleOutlined,
  CloseCircleOutlined,
  SearchOutlined,
  SoundOutlined,
  UnorderedListOutlined,
  FileTextOutlined,
  UpOutlined,
  DownOutlined,
  MenuOutlined,
  EditOutlined,
  DeleteOutlined,
  EyeOutlined
} from '@ant-design/icons-vue'
import { useRouter } from "vue-router";
import dayjs from 'dayjs'; // Import dayjs for date handling

const router = useRouter()
const famousQuotes = ref([])
const famousQuotesIndex = ref(0)

// 搜索框功能
const searchValue = ref('')
const onQuery = () => {
  if (searchValue.value){
    router.push({path:"blog",query:{q:searchValue.value}})
  }else{
    router.push("blog")
  }
}

// 顶部文字显示区域
const announcement = ref('欢迎来到你的个人仪表盘。这是一个展示公告的区域，可以显示重要信息，最多显示两行文字。如果文字太长会被截断，鼠标悬停时会显示完整内容。这是一段测试文字，用来测试文字截断功能和鼠标悬停效果。')
const showTopSection = ref(true)

// 左侧Todo列表
const todos = ref([
])
const todoDones = ref([

])
const newTodo = ref('')
const showLeftSection = ref(true)

// Todo 对话框控制
const todoDialogVisible = ref(false)
const todoDialogMode = ref('add') // 'add' 或 'edit'
const currentTodo = reactive({
  id: null,
  theme: '',
  content: '',
  status: 0,
  type: 1,
  num: 1,
  from: Number(dayjs().unix()), // Use dayjs object directly for date picker
  to: Number(dayjs().add(7, 'day').unix()) // Use dayjs object directly for date picker
})
const todoDetailsVisible = ref(false)
const selectedTodo = ref(null)

const resetTodoForm = () => {
  currentTodo.id = null
  currentTodo.theme = ''
  currentTodo.content = ''
  currentTodo.status = 0
  currentTodo.type = 1
  currentTodo.num = 1
  currentTodo.from = Number(dayjs().unix()) // Use dayjs object directly
  currentTodo.to = Number(dayjs().add(7, 'day').unix()) // Use dayjs object directly
}

const openTodoDialog = (mode, todo = null) => {
  todoDialogMode.value = mode
  resetTodoForm()
  console.log(dayjs().valueOf())
  if (todo) {
    currentTodo.id = todo.id
    currentTodo.theme = todo.theme
    currentTodo.content = todo.content
    currentTodo.status = todo.status
    currentTodo.type = todo.type
    currentTodo.num = todo.num || 1
    currentTodo.from = todo.from ? todo.from : Number(dayjs().unix()) // Convert to dayjs
    currentTodo.to = todo.to ? todo.to : Number(dayjs().add(7, 'day').unix()) // Convert to dayjs
  }

  todoDialogVisible.value = true
}

const closeTodoDialog = () => {
  todoDialogVisible.value = false
}

const saveTodo = () => {
  if (!currentTodo.theme.trim()) return
  todoSave(currentTodo).then(res=>{
    if(res&&res.code===200){
      closeTodoDialog()
    }
  }).finally(()=>{
    fetchTodoList()
  })

}

const showTodoDetails = (todo) => {
  selectedTodo.value = todo
  todoDetailsVisible.value = true
}

const closeTodoDetails = () => {
  todoDetailsVisible.value = false
  selectedTodo.value = null
}

const toggleTodoStatus = (todo) => {
  if(todo.status === 1){
    return
  }
  todoDone({id:todo.id}).then(res=>{
    if(res&&res.code===200){
      fetchTodo()
    }
  })
  // todo.status = todo.status === 0 ? 1 : 0
}

const deleteTodo = (id) => {
  // todos.value = todos.value.filter(todo => todo.id !== id)
  todoDelete({id:id}).then(res=>{
    if(res&&res.code===200){
      fetchTodoList()
    }
  }).finally(()=>{

  })
}
const deleteDoneTodo = (id) => {
  // todos.value = todos.value.filter(todo => todo.id !== id)
  todoDoneDelete({id:id}).then(res=>{
    if(res&&res.code===200){
      fetchTodoDoneList()
    }
  })
}
const pendingTodos = ref([])

const completedTodos = ref([])

// 右侧备忘录
const memos = ref([])
const newMemo = ref('')
const newMemoName = ref('')
const newMemoFamiliarity = ref(3)
const showRightSection = ref(true)

// 备忘录对话框控制
const memoDialogVisible = ref(false)
const memoDialogMode = ref('add') // 'add' 或 'edit'
const currentMemo = reactive({
  id: null,
  name: '',
  content: '',
  status: 3
})
const memoDetailsVisible = ref(false)
const selectedMemo = ref(null)
const statusDialogVisible = ref(false)
const editingFamiliarityMemoId = ref(null)

const resetMemoForm = () => {
  currentMemo.id = null
  currentMemo.name = ''
  currentMemo.content = ''
  currentMemo.status = 3
}

const openMemoDialog = (mode, memo = null) => {
  memoDialogMode.value = mode
  resetMemoForm()

  if (memo) {
    currentMemo.id = memo.id
    currentMemo.name = memo.name
    currentMemo.content = memo.content
    currentMemo.status = memo.status
  }

  memoDialogVisible.value = true
}

const closeMemoDialog = () => {
  memoDialogVisible.value = false
}

const saveMemo = () => {
  if (!currentMemo.name.trim()) return
  // if (memoDialogMode.value === 'add') {
  //   memos.value.push({
  //     id: Date.now(),
  //     name: currentMemo.name,
  //     content: currentMemo.content,
  //     status: currentMemo.status
  //   })
  // } else if (memoDialogMode.value === 'edit') {
  //   const index = memos.value.findIndex(m => m.id === currentMemo.id)
  //   if (index !== -1) {
  //     memos.value[index] = {
  //       ...memos.value[index],
  //       name: currentMemo.name,
  //       content: currentMemo.content,
  //       status: currentMemo.status
  //     }
  //   }
  // }
  memoSave(currentMemo).then(res=>{
    if(res&&res.code===200){
      closeMemoDialog()
    }
  }).finally(()=>{
    fetchMemoList()
  })
}

const showMemoDetails = (memo) => {
  selectedMemo.value = memo
  memoDetailsVisible.value = true
}

const closeMemoDetails = () => {
  memoDetailsVisible.value = false
  selectedMemo.value = null
}

const openFamiliarityDialog = (memoId) => {
  editingFamiliarityMemoId.value = memoId
  statusDialogVisible.value = true
}

const closeFamiliarityDialog = () => {
  statusDialogVisible.value = false
  editingFamiliarityMemoId.value = null
}

const updateMemoFamiliarity = (level) => {
  // const index = memos.value.findIndex(m => m.id === editingFamiliarityMemoId.value)
  // if (index !== -1) {
  //   memos.value[index].status = level
  // }
  memoDone({id:editingFamiliarityMemoId.value,state:level}).then(res=>{
    closeFamiliarityDialog()
  }).finally(()=>{
    fetchMemoList()
  })
}

const deleteMemo = (id) => {
  memoDelete({id:id}).then(() => {

  }).finally(()=>{
    fetchMemoList()
  })
}

const statusColor = (level) => {
  const colors = {
    1: 'bg-red-200/70',
    2: 'bg-orange-200/70',
    3: 'bg-yellow-200/70',
    4: 'bg-green-200/70',
    5: 'bg-emerald-200/70'
  }
  return colors[level] || 'bg-gray-200/70'
}

// 底部音乐播放器
const audioSrc = ref('https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3')
const audioTitle = ref('平静湖水 - 自然之声')
const audioLyrics = ref(`清晨的湖面\n雾气轻笼\n小舟穿行\n引来飞鸟一阵惊鸿`)
const isPlaying = ref(false)
const showBottomSection = ref(true)

const audioPlayer = ref(null)

const togglePlay = () => {
  if (audioPlayer.value) {
    if (isPlaying.value) {
      audioPlayer.value.pause()
    } else {
      audioPlayer.value.play()
    }
    isPlaying.value = !isPlaying.value
  }
}

// 媒体查询检测
const isMobile = ref(false)
const mobileMenuOpen = ref(false)

// 更新视口尺寸
const updateViewportSize = () => {
  isMobile.value = window.innerWidth < 768

  // 在移动端初次加载时，默认显示所有部分
  if (isMobile.value && window.innerWidth < 480) {
    showLeftSection.value = true
    showRightSection.value = true
  }
}

const changeQuote = () => {
  if(famousQuotesIndex.value >= famousQuotes.value.length - 1){
    fetchMessages()
  } else {
    famousQuotesIndex.value = famousQuotesIndex.value + 1
  }
}

const fetchMessages = () => {
  messages({pageSize:10}).then(res => {
    if(res && res.code === 200){
      famousQuotes.value = res.data && res.data.data || []
    }
  }).finally(() => {
    famousQuotesIndex.value = 0
  })
}

const fetchMemoList = () => {
  memoList({page:{pageSize:15,pageNum:1}}).then(res => {
    if(res && res.code === 200){
      memos.value = res.data && res.data.list || []
    }
  })
}
const fetchTodoList = ()=>{
  todoList({page:{pageSize:15,pageNum:1}}).then(res=>{
    if(res && res.code === 200){
      const list = res.data&&res.data.list || []
      list.forEach((item)=>{
        item.status = 0
      })
      pendingTodos.value = list
    }
  })
}
const fetchTodoDoneList = ()=>{
  todoDoneList({page:{pageSize:15,pageNum:1}}).then(res=>{
    if(res && res.code === 200){
      const list = res.data&&res.data.list || []
      completedTodos.value = list
    }
  })
}

const fetchTodo =()=>{
  fetchTodoList()
  fetchTodoDoneList()
}
onMounted(() => {
  fetchMessages()
  fetchMemoList()
  fetchTodo()


  updateViewportSize()
  window.addEventListener('resize', updateViewportSize)
})

onUnmounted(() => {
  window.removeEventListener('resize', updateViewportSize)
})

// 折叠/展开功能
const toggleSection = (section) => {
  switch(section) {
    case 'top':
      showTopSection.value = !showTopSection.value
      break
    case 'left':
      showLeftSection.value = !showLeftSection.value
      break
    case 'right':
      showRightSection.value = !showRightSection.value
      break
    case 'bottom':
      showBottomSection.value = !showBottomSection.value
      break
  }
}

// 计算PC端公告和音乐播放器收缩按钮的位置
const bottomButtonsPosition = computed(() => {
  // 两个都收缩时，对称布局
  if (!showTopSection.value && !showBottomSection.value) {
    return {
      topButton: { left: 'calc(50% - 36px)' },
      bottomButton: { left: 'calc(50% + 36px)' }
    }
  }
  // 只有一个收缩时，居中显示
  if (!showTopSection.value) {
    return {
      topButton: { left: 'calc(50% - 24px)' },
      bottomButton: { left: '50%' }
    }
  }
  if (!showBottomSection.value) {
    return {
      topButton: { left: '50%' },
      bottomButton: { left: 'calc(50% + 24px)' }
    }
  }
  // 默认位置
  return {
    topButton: { left: '50%' },
    bottomButton: { left: '50%', marginLeft: '4rem' }
  }
})

// 移动端菜单切换
const toggleMobileMenu = () => {
  mobileMenuOpen.value = !mobileMenuOpen.value
}

// 计算移动端搜索框的位置和间距
const searchBoxSpacing = computed(() => {
  if (!isMobile.value) return {}

  // 根据可见区块数量调整搜索框位置
  return {
    marginTop: showTopSection.value ? '1.5rem' : '4rem',
    marginBottom: '1.5rem'
  }
})

// 计算底部按钮的布局
const mobileBottomButtons = computed(() => {
  // 计算有多少按钮需要显示
  const buttons = []

  if (!showTopSection.value) buttons.push('top')
  if (!showBottomSection.value) buttons.push('bottom')
  if (!showLeftSection.value) buttons.push('left')
  if (!showRightSection.value) buttons.push('right')

  const totalButtons = buttons.length

  // 根据按钮数量计算每个按钮的位置
  const positions = {}

  if (totalButtons === 1) {
    // 只有一个按钮时居中
    positions[buttons[0]] = { left: '50%' }
  } else if (totalButtons === 2) {
    // 两个按钮时，均匀分布
    positions[buttons[0]] = { left: '33%' }
    positions[buttons[1]] = { left: '67%' }
  } else if (totalButtons === 3) {
    // 三个按钮时，均匀分布
    positions[buttons[0]] = { left: '25%' }
    positions[buttons[1]] = { left: '50%' }
    positions[buttons[2]] = { left: '75%' }
  } else if (totalButtons === 4) {
    // 四个按钮时，均匀分布
    positions[buttons[0]] = { left: '20%' }
    positions[buttons[1]] = { left: '40%' }
    positions[buttons[2]] = { left: '60%' }
    positions[buttons[3]] = { left: '80%' }
  }

  return positions
})

// 创建隐藏区域的占位符样式
const hiddenSectionPlaceholder = computed(() => {
  if (!isMobile.value) return {}

  const styles = {}

  if (!showLeftSection.value) {
    styles.todoPlaceholder = { height: '2rem', marginBottom: '1rem' }
  }

  if (!showRightSection.value) {
    styles.memoPlaceholder = { height: '2rem', marginBottom: '1rem' }
  }

  if (!showBottomSection.value) {
    styles.musicPlaceholder = { height: '2rem', marginBottom: '1rem' }
  }

  return styles
})
</script>

<template>
  <div class="h-full w-full overflow-auto bg-transparent text-gray-800 relative md:my-6">
    <!-- 移动端菜单按钮 -->
    <button v-if="isMobile"
            @click="toggleMobileMenu"
            class="fixed z-20 top-4 right-4 bg-white/50 w-10 h-10 rounded-full flex items-center justify-center shadow-md hover:bg-gray-100/50 text-gray-600">
      <MenuOutlined />
    </button>

    <!-- 移动端导航菜单 -->
    <div v-if="isMobile && mobileMenuOpen"
         class="fixed inset-0 bg-black/30 backdrop-blur-sm z-30 flex items-center justify-center"
         @click="mobileMenuOpen = false">
      <div class="bg-white/90 rounded-lg p-4 w-64" @click.stop>
        <div class="flex flex-col space-y-4">
          <button @click="() => { toggleSection('left'); mobileMenuOpen = false; }"
                  class="flex items-center space-x-2 p-2 hover:bg-gray-100 rounded-md">
            <UnorderedListOutlined />
            <span>{{ showLeftSection ? '隐藏' : '显示' }}待办事项</span>
          </button>
          <button @click="() => { toggleSection('right'); mobileMenuOpen = false; }"
                  class="flex items-center space-x-2 p-2 hover:bg-gray-100 rounded-md">
            <FileTextOutlined />
            <span>{{ showRightSection ? '隐藏' : '显示' }}备忘录</span>
          </button>
          <button @click="() => { toggleSection('top'); mobileMenuOpen = false; }"
                  class="flex items-center space-x-2 p-2 hover:bg-gray-100 rounded-md">
            <UpOutlined v-if="showTopSection" />
            <DownOutlined v-else />
            <span>{{ showTopSection ? '隐藏' : '显示' }}公告</span>
          </button>
          <button @click="() => { toggleSection('bottom'); mobileMenuOpen = false; }"
                  class="flex items-center space-x-2 p-2 hover:bg-gray-100 rounded-md">
            <SoundOutlined />
            <span>{{ showBottomSection ? '隐藏' : '显示' }}音乐播放器</span>
          </button>
        </div>
      </div>
    </div>

    <!-- 主布局容器 -->
    <div class="max-w-full mx-auto grid grid-cols-1 md:grid-cols-4 gap-4 min-h-screen">
      <!-- PC布局 -->
      <template v-if="!isMobile">
        <!-- 左侧区域 - 待办事项 -->
        <div class="md:col-span-1 px-2 md:px-12 h-full flex flex-col">
          <div v-if="showLeftSection" class="bg-white/30 backdrop-blur-sm p-4 rounded-lg shadow-sm border border-gray-200/50 flex-1 overflow-y-auto">
            <div class="flex justify-between items-center mb-4">
              <h2 class="text-lg font-semibold text-gray-700 flex items-center">
                <UnorderedListOutlined class="mr-2" />待办事项
              </h2>
              <div class="flex items-center">
                <button @click="openTodoDialog('add')" class="text-gray-500 hover:text-green-600 mr-2">
                  <PlusCircleOutlined />
                </button>
                <button @click="toggleSection('left')" class="text-gray-500 hover:text-gray-700">
                  <MinusCircleOutlined />
                </button>
              </div>
            </div>

            <h3 class="font-medium text-gray-600 mb-2">待完成</h3>
            <ul class="space-y-2 mb-4">
              <li v-for="todo in pendingTodos" :key="todo.id"
                  class="flex items-center p-2 hover:bg-gray-50/30 rounded-md transition-colors">
                <button @click="toggleTodoStatus(todo)" class="flex-shrink-0 mr-2 text-gray-400 hover:text-green-500">
                  <PlusCircleOutlined />
                </button>
                <div class="flex-1 flex flex-col overflow-hidden">
                  <button @click="showTodoDetails(todo)" class="text-left truncate hover:text-green-600">
                    {{ todo.theme }}
                  </button>
                </div>
                <div class="flex items-center">
                  <button @click="openTodoDialog('edit', todo)" class="text-gray-400 hover:text-blue-500 mr-2">
                    <EditOutlined />
                  </button>
                  <button @click="deleteTodo(todo.id)" class="text-gray-400 hover:text-red-500">
                    <DeleteOutlined />
                  </button>
                </div>
              </li>
              <li v-if="pendingTodos.length === 0" class="text-gray-400 italic text-sm">暂无待办事项</li>
            </ul>

            <h3 class="font-medium text-gray-600 mb-2">已完成</h3>
            <ul class="space-y-2">
              <li v-for="todo in completedTodos" :key="todo.id"
                  class="flex items-center p-2 hover:bg-gray-50/30 rounded-md transition-colors">
                <button @click="toggleTodoStatus(todo)" class="flex-shrink-0 mr-2 text-green-500 hover:text-green-600">
                  <CheckCircleOutlined />
                </button>
                <div class="flex-1 flex flex-col overflow-hidden">
                  <button @click="showTodoDetails(todo)" class="text-left truncate line-through text-gray-400 hover:text-green-600">
                    {{ todo.theme }}
                  </button>
                </div>
                <button @click="deleteDoneTodo(todo.id)" class="text-gray-400 hover:text-red-500">
                  <DeleteOutlined />
                </button>
              </li>
              <li v-if="completedTodos.length === 0" class="text-gray-400 italic text-sm">暂无已完成事项</li>
            </ul>
          </div>
        </div>

        <!-- 中间区域 - 分为上中下三个部分 -->
        <div class="md:col-span-2 px-2 md:px-12 flex flex-col h-full">
          <!-- 上部分 - 公告区域 -->
          <div class="md:h-[15%] mb-1">
            <div v-if="showTopSection&&famousQuotes.length>0" class="bg-transparent p-4 rounded-lg h-full relative" @click="changeQuote">
              <p class="line-clamp-2 text-gray-700 group-hover:line-clamp-none relative group text-center cursor-pointer">
                {{ famousQuotes[famousQuotesIndex].content }}
                <span v-if="false" class="absolute inset-0 opacity-0 group-hover:opacity-100 bg-white/50 backdrop-blur-sm p-4 rounded-lg shadow-md z-10 transition-opacity duration-300">
                  {{ famousQuotes[famousQuotesIndex].content }}
                </span>
              </p>
              <button v-if="false" @click="toggleSection('top')" class="absolute right-2 top-2 text-gray-500 hover:text-gray-700 z-20">
                <UpOutlined />
              </button>
            </div>
          </div>

          <!-- 中部分 - 搜索框 -->
          <div class="md:h-[30%] flex items-center justify-center mb-4">
            <div class="w-full">
              <div class="flex">
                <input v-model="searchValue"
                       class="flex-1 px-4 py-3 h-12 rounded-l-lg border-0 outline-none focus:ring-2 focus:ring-green-200 bg-gradient-to-r from-green-100/60 to-gray-100/60 placeholder-gray-500"
                       placeholder="you can search blog title from here!" />
                <button @click="onQuery"
                        class="px-6 h-12 bg-white/50 hover:bg-white/70 rounded-r-lg text-gray-700 shadow-sm border border-gray-200/50 transition-colors">
                  <SearchOutlined />
                </button>
              </div>
            </div>
          </div>

          <!-- 下部分 - 音乐播放器 -->
          <div class="flex-1">
            <div v-if="showBottomSection" class="bg-white/30 backdrop-blur-sm p-4 rounded-lg shadow-sm border border-gray-200/50 h-full">
              <div class="flex justify-between items-center mb-2">
                <h2 class="text-lg font-semibold text-gray-700 flex items-center">
                  <SoundOutlined class="mr-2" />音乐播放器
                </h2>
                <button @click="toggleSection('bottom')" class="text-gray-500 hover:text-gray-700">
                  <MinusCircleOutlined />
                </button>
              </div>

              <audio ref="audioPlayer" :src="audioSrc" class="hidden"></audio>

              <div class="flex flex-col md:flex-row gap-4 h-[calc(100%-2rem)]">
                <div class="flex-1">
                  <div class="flex items-center space-x-3 mb-3">
                    <button @click="togglePlay"
                            class="w-10 h-10 rounded-full bg-green-100/50 hover:bg-green-200/50 flex items-center justify-center text-gray-700">
                      <span v-if="isPlaying">⏸️</span>
                      <span v-else>▶️</span>
                    </button>
                    <div>
                      <div class="font-medium">{{ audioTitle }}</div>
                      <div class="text-sm text-gray-500">自然音乐</div>
                    </div>
                  </div>

                  <div class="w-full bg-gray-200/50 rounded-full h-1.5 mb-4">
                    <div class="bg-green-300/70 h-1.5 rounded-full w-1/3"></div>
                  </div>
                </div>

                <div class="flex-1 bg-gray-50/30 p-3 rounded-md overflow-y-auto h-full">
                  <p class="text-sm text-gray-600 whitespace-pre-line">{{ audioLyrics }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 右侧区域 - 备忘录 -->
        <div class="md:col-span-1 px-2 md:px-12 h-full flex flex-col">
          <div v-if="showRightSection" class="bg-white/30 backdrop-blur-sm p-4 rounded-lg shadow-sm border border-gray-200/50 flex-1 overflow-y-auto">
            <div class="flex justify-between items-center mb-4">
              <h2 class="text-lg font-semibold text-gray-700 flex items-center">
                <FileTextOutlined class="mr-2" />备忘录
              </h2>
              <div class="flex items-center">
                <button @click="openMemoDialog('add')" class="text-gray-500 hover:text-green-600 mr-2">
                  <PlusCircleOutlined />
                </button>
                <button @click="toggleSection('right')" class="text-gray-500 hover:text-gray-700">
                  <MinusCircleOutlined />
                </button>
              </div>
            </div>

            <ul class="space-y-3">
              <li v-for="memo in memos" :key="memo.id"
                  :class="[
                    'relative p-3 rounded-md border border-l-4 grid grid-cols-[1fr_auto]',
                    statusColor(memo.status),
                    'border-l-' + (memo.status >= 4 ? 'green-400' : memo.status >= 3 ? 'yellow-400' : 'orange-400')
                  ]">
                <div class="overflow-hidden pr-2">
                  <div @click="showMemoDetails(memo)" class="cursor-pointer">
                    <h3 class="font-medium truncate hover:text-green-600">{{ memo.name }}</h3>
                    <p class="text-sm text-gray-600 line-clamp-2">{{ memo.content }}</p>
                  </div>
                </div>
                <div class="flex flex-col items-center justify-between h-full">
                  <div class="flex flex-col gap-3 mb-2">
                    <button @click="openMemoDialog('edit', memo)" class="text-gray-400 hover:text-blue-500">
                      <EditOutlined />
                    </button>
                    <button @click="deleteMemo(memo.id)" class="text-gray-400 hover:text-red-500">
                      <DeleteOutlined />
                    </button>
                  </div>
                  <button @click="openFamiliarityDialog(memo.id)" class="text-xs font-medium px-1.5 py-0.5 rounded-full bg-white/30 hover:bg-white/50">
                    {{ memo.status }}/5
                  </button>
                </div>
              </li>
              <li v-if="memos.length === 0" class="text-gray-400 italic text-sm">暂无备忘录</li>
            </ul>
          </div>
        </div>
      </template>

      <!-- 移动端布局 - 垂直堆叠 -->
      <template v-else>
        <div class="flex flex-col min-h-screen justify-between pb-16">
          <!-- 顶部区域 -->
          <div class="flex-grow">
            <!-- 公告区域 - 即使隐藏也保留空间 -->
            <div class="relative mb-4" v-if="famousQuotes.length>0">
              <div v-if="showTopSection" class="bg-transparent p-4 rounded-lg"  @click="changeQuote">
                <p class="line-clamp-2  text-gray-700 group-hover:line-clamp-none relative group text-center cursor-pointer">
                  {{ famousQuotes[famousQuotesIndex].content }}
                  <span v-if="false" class="absolute inset-0 opacity-0 group-hover:opacity-100 bg-white/50 backdrop-blur-sm p-4 rounded-lg shadow-md z-10 transition-opacity duration-300">
                    {{ famousQuotes[famousQuotesIndex].content }}
                  </span>
                </p>
                <button v-if="false" @click="toggleSection('top')" class="absolute right-2 top-2 text-gray-500 hover:text-gray-700 z-20">
                  <UpOutlined />
                </button>
              </div>
              <!-- 隐藏时的占位符 -->
              <div v-else class="h-10 flex items-center justify-center">
                <button @click="toggleSection('top')" class="bg-white/30 backdrop-blur-sm px-4 py-2 rounded-lg text-gray-600 flex items-center">
                  <DownOutlined class="mr-2" />
                  <span>显示公告</span>
                </button>
              </div>
            </div>

            <!-- 搜索框 - 不使用fixed定位 -->
            <div :style="searchBoxSpacing" class="transition-all duration-300 bg-white/30 backdrop-blur-sm p-3 rounded-lg shadow-sm mb-4">
              <div class="flex">
                <input v-model="searchValue"
                       class="flex-1 px-4 py-3 h-12 rounded-l-lg border-0 outline-none focus:ring-2 focus:ring-green-200 bg-gradient-to-r from-green-100/60 to-gray-100/60 placeholder-gray-500"
                       placeholder="you can search blog title from here!" />
                <button @click="onQuery"
                        class="px-6 h-12 bg-white/50 hover:bg-white/70 rounded-r-lg text-gray-700 shadow-sm border border-gray-200/50 transition-colors">
                  <SearchOutlined />
                </button>
              </div>
            </div>

            <!-- 待办事项 - 隐藏时保留占位符 -->
            <div v-if="showLeftSection" class="bg-white/30 backdrop-blur-sm p-4 rounded-lg shadow-sm border border-gray-200/50 my-4">
              <div class="flex justify-between items-center mb-4">
                <h2 class="text-lg font-semibold text-gray-700 flex items-center">
                  <UnorderedListOutlined class="mr-2" />待办事项
                </h2>
                <div class="flex items-center">
                  <button @click="openTodoDialog('add')" class="text-gray-500 hover:text-green-600 mr-2">
                    <PlusCircleOutlined />
                  </button>
                  <button @click="toggleSection('left')" class="text-gray-500 hover:text-gray-700">
                    <MinusCircleOutlined />
                  </button>
                </div>
              </div>

              <h3 class="font-medium text-gray-600 mb-2">待完成</h3>
              <ul class="space-y-2 mb-4">
                <li v-for="todo in pendingTodos" :key="todo.id"
                    class="flex items-center p-2 hover:bg-gray-50/30 rounded-md transition-colors">
                  <button @click="toggleTodoStatus(todo)" class="flex-shrink-0 mr-2 text-gray-400 hover:text-green-500">
                    <PlusCircleOutlined />
                  </button>
                  <div class="flex-1 flex flex-col overflow-hidden">
                    <button @click="showTodoDetails(todo)" class="text-left truncate hover:text-green-600">
                      {{ todo.theme }}
                    </button>
                  </div>
                  <div class="flex items-center">
                    <button @click="openTodoDialog('edit', todo)" class="text-gray-400 hover:text-blue-500 mr-2">
                      <EditOutlined />
                    </button>
                    <button @click="deleteTodo(todo.id)" class="text-gray-400 hover:text-red-500">
                      <DeleteOutlined />
                    </button>
                  </div>
                </li>
                <li v-if="pendingTodos.length === 0" class="text-gray-400 italic text-sm">暂无待办事项</li>
              </ul>

              <h3 class="font-medium text-gray-600 mb-2">已完成</h3>
              <ul class="space-y-2">
                <li v-for="todo in completedTodos" :key="todo.id"
                    class="flex items-center p-2 hover:bg-gray-50/30 rounded-md transition-colors">
                  <button @click="toggleTodoStatus(todo)" class="flex-shrink-0 mr-2 text-green-500 hover:text-green-600">
                    <CheckCircleOutlined />
                  </button>
                  <div class="flex-1 flex flex-col overflow-hidden">
                    <button @click="showTodoDetails(todo)" class="text-left truncate line-through text-gray-400 hover:text-green-600">
                      {{ todo.theme }}
                    </button>
                  </div>
                  <button @click="deleteDoneTodo(todo.doneId)" class="text-gray-400 hover:text-red-500">
                    <DeleteOutlined />
                  </button>
                </li>
                <li v-if="completedTodos.length === 0" class="text-gray-400 italic text-sm">暂无已完成事项</li>
              </ul>
            </div>
            <!-- 待办事项隐藏时的占位符 -->
            <div v-else class="my-4" :style="hiddenSectionPlaceholder.todoPlaceholder"></div>

            <!-- 备忘录 - 隐藏时保留占位符 -->
            <div v-if="showRightSection" class="bg-white/30 backdrop-blur-sm p-4 rounded-lg shadow-sm border border-gray-200/50 my-4">
              <div class="flex justify-between items-center mb-4">
                <h2 class="text-lg font-semibold text-gray-700 flex items-center">
                  <FileTextOutlined class="mr-2" />备忘录
                </h2>
                <div class="flex items-center">
                  <button @click="openMemoDialog('add')" class="text-gray-500 hover:text-green-600 mr-2">
                    <PlusCircleOutlined />
                  </button>
                  <button @click="toggleSection('right')" class="text-gray-500 hover:text-gray-700">
                    <MinusCircleOutlined />
                  </button>
                </div>
              </div>

              <ul class="space-y-3">
                <li v-for="memo in memos" :key="memo.id"
                    :class="[
                    'relative p-3 rounded-md border border-l-4 grid grid-cols-[1fr_auto]',
                    statusColor(memo.status),
                    'border-l-' + (memo.status >= 4 ? 'green-400' : memo.status >= 3 ? 'yellow-400' : 'orange-400')
                  ]">
                  <div class="overflow-hidden pr-2">
                    <div @click="showMemoDetails(memo)" class="cursor-pointer">
                      <h3 class="font-medium truncate hover:text-green-600">{{ memo.name }}</h3>
                      <p class="text-sm text-gray-600 line-clamp-2">{{ memo.content }}</p>
                    </div>
                  </div>
                  <div class="flex flex-col items-center justify-between h-full">
                    <div class="flex flex-col gap-3 mb-2">
                      <button @click="openMemoDialog('edit', memo)" class="text-gray-400 hover:text-blue-500">
                        <EditOutlined />
                      </button>
                      <button @click="deleteMemo(memo.id)" class="text-gray-400 hover:text-red-500">
                        <DeleteOutlined />
                      </button>
                    </div>
                    <button @click="openFamiliarityDialog(memo.id)" class="text-xs font-medium px-1.5 py-0.5 rounded-full bg-white/30 hover:bg-white/50">
                      {{ memo.status }}/5
                    </button>
                  </div>
                </li>
                <li v-if="memos.length === 0" class="text-gray-400 italic text-sm">暂无备忘录</li>
              </ul>
            </div>
            <!-- 备忘录隐藏时的占位符 -->
            <div v-else class="my-4" :style="hiddenSectionPlaceholder.memoPlaceholder"></div>

            <!-- 音乐播放器 - 隐藏时保留占位符 -->
            <div v-if="showBottomSection" class="bg-white/30 backdrop-blur-sm p-4 rounded-lg shadow-sm border border-gray-200/50 my-4">
              <div class="flex justify-between items-center mb-2">
                <h2 class="text-lg font-semibold text-gray-700 flex items-center">
                  <SoundOutlined class="mr-2" />音乐播放器
                </h2>
                <button @click="toggleSection('bottom')" class="text-gray-500 hover:text-gray-700">
                  <MinusCircleOutlined />
                </button>
              </div>

              <audio ref="audioPlayer" :src="audioSrc" class="hidden"></audio>

              <div class="flex flex-col gap-4">
                <div class="flex-1">
                  <div class="flex items-center space-x-3 mb-3">
                    <button @click="togglePlay"
                            class="w-10 h-10 rounded-full bg-green-100/50 hover:bg-green-200/50 flex items-center justify-center text-gray-700">
                      <span v-if="isPlaying">⏸️</span>
                      <span v-else>▶️</span>
                    </button>
                    <div>
                      <div class="font-medium">{{ audioTitle }}</div>
                      <div class="text-sm text-gray-500">自然音乐</div>
                    </div>
                  </div>

                  <div class="w-full bg-gray-200/50 rounded-full h-1.5 mb-4">
                    <div class="bg-green-300/70 h-1.5 rounded-full w-1/3"></div>
                  </div>
                </div>

                <div class="flex-1 bg-gray-50/30 p-3 rounded-md overflow-y-auto max-h-32">
                  <p class="text-sm text-gray-600 whitespace-pre-line">{{ audioLyrics }}</p>
                </div>
              </div>
            </div>
            <!-- 音乐播放器隐藏时的占位符 -->
            <div v-else class="my-4" :style="hiddenSectionPlaceholder.musicPlaceholder"></div>
          </div>

          <!-- 底部按钮区域 - 均匀分布按钮 -->
          <div class="fixed bottom-4 left-0 right-0 z-10">
            <div class="relative h-10 flex justify-center">
              <!-- 使用绝对定位和计算的位置来放置按钮 -->
              <button v-if="!showTopSection"
                      @click="toggleSection('top')"
                      class="absolute bg-white/50 w-10 h-10 rounded-full flex items-center justify-center shadow-md hover:bg-gray-100/50 text-gray-600 transform -translate-x-1/2"
                      :style="{ left: mobileBottomButtons.top?.left || '25%' }">
                <DownOutlined />
              </button>

              <button v-if="!showBottomSection"
                      @click="toggleSection('bottom')"
                      class="absolute bg-white/50 w-10 h-10 rounded-full flex items-center justify-center shadow-md hover:bg-gray-100/50 text-gray-600 transform -translate-x-1/2"
                      :style="{ left: mobileBottomButtons.bottom?.left || '50%' }">
                <SoundOutlined />
              </button>

              <button v-if="!showLeftSection"
                      @click="toggleSection('left')"
                      class="absolute bg-white/50 w-10 h-10 rounded-full flex items-center justify-center shadow-md hover:bg-gray-100/50 text-gray-600 transform -translate-x-1/2"
                      :style="{ left: mobileBottomButtons.left?.left || '75%' }">
                <UnorderedListOutlined />
              </button>

              <button v-if="!showRightSection"
                      @click="toggleSection('right')"
                      class="absolute bg-white/50 w-10 h-10 rounded-full flex items-center justify-center shadow-md hover:bg-gray-100/50 text-gray-600 transform -translate-x-1/2"
                      :style="{ left: mobileBottomButtons.right?.left || '100%' }">
                <FileTextOutlined />
              </button>
            </div>
          </div>
        </div>
      </template>
    </div>

    <!-- PC端固定在视口中的折叠/展开按钮 -->
    <button v-if="!showLeftSection && !isMobile"
            @click="toggleSection('left')"
            class="fixed z-10 bg-white/50 w-10 h-10 rounded-full flex items-center justify-center shadow-md hover:bg-gray-100/50 text-gray-600 left-4 top-1/2 transform -translate-y-1/2">
      <UnorderedListOutlined />
    </button>

    <button v-if="!showRightSection && !isMobile"
            @click="toggleSection('right')"
            class="fixed z-10 bg-white/50 w-10 h-10 rounded-full flex items-center justify-center shadow-md hover:bg-gray-100/50 text-gray-600 right-4 top-1/2 transform -translate-y-1/2">
      <FileTextOutlined />
    </button>

    <button v-if="!showTopSection && !isMobile"
            @click="toggleSection('top')"
            class="fixed z-10 bg-white/50 w-10 h-10 rounded-full flex items-center justify-center shadow-md hover:bg-gray-100/50 text-gray-600 bottom-4 transform -translate-x-1/2"
            :style="bottomButtonsPosition.topButton">
      <DownOutlined />
    </button>

    <button v-if="!showBottomSection && !isMobile"
            @click="toggleSection('bottom')"
            class="fixed z-10 bg-white/50 w-10 h-10 rounded-full flex items-center justify-center shadow-md hover:bg-gray-100/50 text-gray-600 bottom-4 transform -translate-x-1/2"
            :style="bottomButtonsPosition.bottomButton">
      <SoundOutlined />
    </button>

    <!-- 备忘录对话框 -->
    <div v-if="memoDialogVisible" class="fixed inset-0 bg-black/30 backdrop-blur-sm z-40 flex items-center justify-center" @click="closeMemoDialog">
      <div class="bg-white/95 rounded-lg p-6 w-full max-w-md shadow-xl" @click.stop>
        <h3 class="text-lg font-semibold mb-4">{{ memoDialogMode === 'add' ? '新增备忘录' : '编辑备忘录' }}</h3>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">标题</label>
            <input
                v-model="currentMemo.name"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-green-200 focus:border-green-300"
                placeholder="请输入标题"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">内容</label>
            <textarea
                v-model="currentMemo.content"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-green-200 focus:border-green-300 min-h-[100px]"
                placeholder="请输入内容"
            ></textarea>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">熟悉度</label>
            <div class="flex items-center space-x-2">
              <button v-for="n in 5" :key="n" @click="currentMemo.status = n"
                      :class="[
                  'w-8 h-8 rounded-full text-sm flex items-center justify-center transition-colors',
                  currentMemo.status >= n ? statusColor(n) : 'bg-gray-100'
                ]">
                {{ n }}
              </button>
            </div>
          </div>
        </div>

        <div class="flex justify-end space-x-2 mt-6">
          <button @click="closeMemoDialog" class="px-4 py-2 border border-gray-300 rounded-md hover:bg-gray-50">
            取消
          </button>
          <button @click="saveMemo" class="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600">
            确定
          </button>
        </div>
      </div>
    </div>

    <!-- 备忘录详情对话框 -->
    <div v-if="memoDetailsVisible" class="fixed inset-0 bg-black/30 backdrop-blur-sm z-40 flex items-center justify-center" @click="closeMemoDetails">
      <div class="bg-white/95 rounded-lg p-6 w-full max-w-md shadow-xl" @click.stop>
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-semibold">备忘录详情</h3>
          <button @click="closeMemoDetails" class="text-gray-500 hover:text-gray-700">
            <CloseCircleOutlined />
          </button>
        </div>

        <div v-if="selectedMemo" class="space-y-4">
          <div>
            <h4 class="text-sm font-medium text-gray-500">标题</h4>
            <p class="mt-1">{{ selectedMemo.name }}</p>
          </div>

          <div>
            <h4 class="text-sm font-medium text-gray-500">内容</h4>
            <p class="mt-1 whitespace-pre-line break-words max-w-full overflow-hidden">{{ selectedMemo.content }}</p>
          </div>

          <div>
            <h4 class="text-sm font-medium text-gray-500">熟悉度</h4>
            <div class="mt-1 flex items-center space-x-1">
              <div v-for="n in 5" :key="n"
                   :class="[
                  'w-6 h-6 rounded-full text-xs flex items-center justify-center',
                  selectedMemo.status >= n ? statusColor(n) : 'bg-gray-100/50'
                ]">
                {{ n }}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 熟悉度选择对话框 -->
    <div v-if="statusDialogVisible" class="fixed inset-0 bg-black/30 backdrop-blur-sm z-40 flex items-center justify-center" @click="closeFamiliarityDialog">
      <div class="bg-white/95 rounded-lg p-4 shadow-xl" @click.stop>
        <h3 class="text-base font-medium mb-3 text-center">选择熟悉度</h3>
        <div class="flex items-center space-x-2 p-2">
          <button v-for="n in 5" :key="n" @click="updateMemoFamiliarity(n)"
                  :class="[
              'w-10 h-10 rounded-full text-sm flex items-center justify-center hover:scale-110 transition-transform',
              statusColor(n)
            ]">
            {{ n }}
          </button>
        </div>
      </div>
    </div>

    <!-- Todo对话框 -->
    <div v-if="todoDialogVisible" class="fixed inset-0 bg-black/30 backdrop-blur-sm z-40 flex items-center justify-center" @click="closeTodoDialog">
      <div class="bg-white/95 rounded-lg p-6 w-full max-w-md shadow-xl" @click.stop>
        <h3 class="text-lg font-semibold mb-4">{{ todoDialogMode === 'add' ? '新增待办事项' : '编辑待办事项' }}</h3>

        <div class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">标题</label>
            <input
                v-model="currentTodo.theme"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-green-200 focus:border-green-300"
                placeholder="请输入标题"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">内容</label>
            <textarea
                v-model="currentTodo.content"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-green-200 focus:border-green-300 min-h-[80px]"
                placeholder="请输入内容"
            ></textarea>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">类型</label>
            <div class="grid grid-cols-3 gap-2">
              <button
                  @click="currentTodo.type = 1"
                  :class="[
                  'px-3 py-2 border rounded-md text-sm',
                  currentTodo.type === 1 ? 'bg-green-100 border-green-300' : 'border-gray-300 hover:bg-gray-50'
                ]"
              >
                按次
              </button>
              <button
                  @click="currentTodo.type = 2"
                  :class="[
                  'px-3 py-2 border rounded-md text-sm',
                  currentTodo.type === 2 ? 'bg-green-100 border-green-300' : 'border-gray-300 hover:bg-gray-50'
                ]"
              >
                固定日期
              </button>
              <button
                  @click="currentTodo.type = 3"
                  :class="[
                  'px-3 py-2 border rounded-md text-sm',
                  currentTodo.type === 3 ? 'bg-green-100 border-green-300' : 'border-gray-300 hover:bg-gray-50'
                ]"
              >
                无固定日期
              </button>
            </div>
          </div>
          <!-- 按类型显示不同的输入字段 -->
          <div v-if="currentTodo.type === 1">
            <label class="block text-sm font-medium text-gray-700 mb-1">执行次数</label>
            <input
                v-model.number="currentTodo.num"
                type="number"
                min="1"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-green-200 focus:border-green-300"
            />
          </div>

          <div v-if="currentTodo.type === 2" class="space-y-2">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">开始日期</label>
              <a-date-picker
                  v-model:value="currentTodo.from"
                  class="w-full"
                  value-format="X"
                  :placeholder="'请选择开始日期'"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">结束日期</label>
              <a-date-picker
                  v-model:value="currentTodo.to"
                  class="w-full"
                  value-format="X"
                  :placeholder="'请选择结束日期'"
              />
            </div>
          </div>

          <div v-if="currentTodo.type === 3">
            <label class="block text-sm font-medium text-gray-700 mb-1">开始日期</label>
            <a-date-picker
                v-model:value="currentTodo.from"
                class="w-full"
                value-format="X"
                :placeholder="'请选择开始日期'"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">状态</label>
            <div class="grid grid-cols-2 gap-2">
              <button
                  @click="currentTodo.status = 0"
                  :class="[
                  'px-3 py-2 border rounded-md text-sm',
                  currentTodo.status === 0 ? 'bg-green-100 border-green-300' : 'border-gray-300 hover:bg-gray-50'
                ]"
              >
                待办
              </button>
              <button
                  @click="currentTodo.status = 1"
                  :class="[
                  'px-3 py-2 border rounded-md text-sm',
                  currentTodo.status === 1 ? 'bg-green-100 border-green-300' : 'border-gray-300 hover:bg-gray-50'
                ]"
              >
                已做
              </button>
            </div>
          </div>
        </div>

        <div class="flex justify-end space-x-2 mt-6">
          <button @click="closeTodoDialog" class="px-4 py-2 border border-gray-300 rounded-md hover:bg-gray-50">
            取消
          </button>
          <button @click="saveTodo" class="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600">
            确定
          </button>
        </div>
      </div>
    </div>

    <!-- Todo详情对话框 -->
    <div v-if="todoDetailsVisible" class="fixed inset-0 bg-black/30 backdrop-blur-sm z-40 flex items-center justify-center" @click="closeTodoDetails">
      <div class="bg-white/95 rounded-lg p-6 w-full max-w-md shadow-xl" @click.stop>
        <div class="flex justify-between items-center mb-4">
          <h3 class="text-lg font-semibold">待办事项详情</h3>
          <button @click="closeTodoDetails" class="text-gray-500 hover:text-gray-700">
            <CloseCircleOutlined />
          </button>
        </div>

        <div v-if="selectedTodo" class="space-y-4">
          <div>
            <h4 class="text-sm font-medium text-gray-500">标题</h4>
            <p class="mt-1">{{ selectedTodo.theme }}</p>
          </div>

          <div>
            <h4 class="text-sm font-medium text-gray-500">内容</h4>
            <p class="mt-1 whitespace-pre-line break-words max-w-full overflow-hidden">{{ selectedTodo.content }}</p>
          </div>

          <div>
            <h4 class="text-sm font-medium text-gray-500">类型</h4>
            <p class="mt-1">
              {{ selectedTodo.type === 1 ? '按次' : selectedTodo.type === 2 ? '固定日期' : '无固定日期' }}
            </p>
          </div>

          <div v-if="selectedTodo.type === 1">
            <h4 class="text-sm font-medium text-gray-500">执行次数</h4>
            <p class="mt-1">{{ selectedTodo.num || 1 }}</p>
          </div>

          <div v-if="selectedTodo.type === 2 || selectedTodo.type === 3">
            <h4 class="text-sm font-medium text-gray-500">开始日期</h4>
            <p class="mt-1">{{ selectedTodo.from ? $formatDate(selectedTodo.from,'{y}-{m}-{d}'): '未设置' }} </p>
          </div>

          <div v-if="selectedTodo.type === 2">
            <h4 class="text-sm font-medium text-gray-500">结束日期</h4>
            <p class="mt-1">{{ selectedTodo.to ? $formatDate(selectedTodo.to,'{y}-{m}-{d}') : '未设置' }}</p>
          </div>

          <div>
            <h4 class="text-sm font-medium text-gray-500">状态</h4>
            <p class="mt-1">{{ selectedTodo.status === 0 ? '待办' : '已完成' }}</p>
          </div>

          <div>
            <h4 class="text-sm font-medium text-gray-500">创建时间</h4>
            <p class="mt-1">{{ selectedTodo.createdAt ? $formatDate(selectedTodo.createdAt,'{y}-{m}-{d}') : '未知' }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* 添加背景图片 */
:deep(body) {
  background-image: url('https://hebbkx1anhila5yf.public.blob.vercel-storage.com/image-SWgTn3cmLGDHEHltsiuw5oD74JL5Db.png');
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  background-attachment: fixed;
}

/* 移动端样式调整 */
@media (max-width: 768px) {
  .min-h-screen {
    min-height: 100vh;
  }
}
</style>
