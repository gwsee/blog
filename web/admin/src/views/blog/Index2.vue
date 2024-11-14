<template>
  <a-row style="background-color: green;width: 100%"  >
    <a-col :md="5" :sm="24" :xs="24" style="background-color: #1a1a1a" >
      1
    </a-col>
    <a-col :md="14" :sm="24" :xs="24"  style="background-color: #992929;text-align: left" >
      <a-list item-layout="vertical" size="large" :pagination="pagination" :data-source="listData">
        <template #renderItem="{ item }">
          <a-list-item key="item.title">
            <template #actions>
          <span v-for="{ icon, text } in actions" :key="icon">
            <component :is="icon" style="margin-right: 8px" />
            {{ text }}
          </span>
            </template>
            <template #extra>
              <img
                  width="272"
                  alt="logo"
                  src="https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png"
              />
            </template>
            <a-list-item-meta :description="item.description">
              <template #title>
                <a :href="item.href">{{ item.title }}</a>
              </template>
              <template #avatar><a-avatar :src="item.avatar" /></template>
            </a-list-item-meta>
            <div>{{ item.content }}</div>
          </a-list-item>
        </template>
      </a-list>
    </a-col>
    <a-col :md="5" :sm="24" :xs="24"  style="background-color: #1a1a1a" >
      3
    </a-col>
  </a-row>
</template>

<script setup>
import { StarOutlined, LikeOutlined, MessageOutlined } from '@ant-design/icons-vue';
const listData = [];

for (let i = 0; i < 23; i++) {
  listData.push({
    href: 'https://www.antdv.com/',
    title: `ant design vue part ${i}`,
    avatar: 'https://joeschmoe.io/api/v1/random',
    description:
        'Ant Design, a design language for background applications, is refined by Ant UED Team.',
    content:
        'We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently. We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.',
  });
}

const pagination = {
  onChange: (page) => {
    console.log(page);
  },
  pageSize: 3,
};
const actions = [
  { icon: StarOutlined, text: '156' },
  { icon: LikeOutlined, text: '156' },
  { icon: MessageOutlined, text: '2' },
];
</script>

<style scoped>
.blog-content-eclipses{
  flex: 1; /* 右边部分自适应 */
  overflow: hidden; /* 超出部分隐藏 */
  text-overflow: ellipsis; /* 文字溢出显示省略号 */
  white-space: nowrap; /* 文字不换行 */
  background-color: lightgreen; /* 仅为了视觉效果 */
  margin: 15px 15px 15px 0;
}
</style>
