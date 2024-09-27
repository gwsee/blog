<template>
  <a-space direction="vertical" :style="{ width: '100%' }" :size="[0, 48]">
    <a-layout>
      <a-layout-header :style="headerStyle">Header</a-layout-header>
      <a-layout-content :style="contentStyle">Content</a-layout-content>
      <a-layout-footer :style="footerStyle">Footer</a-layout-footer>
    </a-layout>

    <a-layout>
      <a-layout-header :style="headerStyle">Header</a-layout-header>
      <a-layout>
        <a-layout-sider :style="siderStyle">Sider</a-layout-sider>
        <a-layout-content :style="contentStyle">Content</a-layout-content>
      </a-layout>
      <a-layout-footer :style="footerStyle">Footer</a-layout-footer>
    </a-layout>

    <a-layout>
      <a-layout-header :style="headerStyle">Header</a-layout-header>
      <a-layout>
        <a-layout-content :style="contentStyle">Content</a-layout-content>
        <a-layout-sider :style="siderStyle">Sider</a-layout-sider>
      </a-layout>
      <a-layout-footer :style="footerStyle">Footer</a-layout-footer>
    </a-layout>

    <a-layout>
      <a-layout-sider :style="siderStyle">Sider</a-layout-sider>
      <a-layout>
        <a-layout-header :style="headerStyle">Header</a-layout-header>
        <a-layout-content :style="contentStyle">Content</a-layout-content>
        <a-layout-footer :style="footerStyle">Footer</a-layout-footer>
      </a-layout>
    </a-layout>
  </a-space>
</template>

<style scoped>
.content {
  display: flex;
  min-height: 100vh;
  line-height: 1.1;
  text-align: center;
  flex-direction: column;
  justify-content: center;
}

.content h1 {
  font-size: 3.6rem;
  font-weight: 700;
}

.content p {
  font-size: 1.2rem;
  font-weight: 400;
  opacity: 0.5;
}
</style>
