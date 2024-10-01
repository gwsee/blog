<template>
  <a-row style="background-color: green;width: 100%"  >
    <a-col :md="5" :sm="24" :xs="24" style="background-color: #1a1a1a" >
      1
    </a-col>
    <a-col :md="14" :sm="24" :xs="24"  style="background-color: #992929;text-align: center" >
      <a-card title="Default size card" style="width: 100%;text-align: left" :hoverable="true" :bodyStyle="{padding:0}">
        <template #extra><a href="#">Edit</a></template>
        <div style="height: 130px;display: flex">
          <a-avatar :size="100" shape="square" style="margin: 15px;">
            <template #icon><UserOutlined /></template>
          </a-avatar>
          <div class="ellipsis-multiline2">
           &nbsp; &nbsp;  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently. We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.
          </div>
        </div>
        <template #actions>
          <span><HeartOutlined /> 999+</span>
          <LikeOutlined ></LikeOutlined>
          <DislikeOutlined />
        </template>
      </a-card>
      <a-card title="Default size card" style="width: 100%;text-align: left" :hoverable="true" :bodyStyle="{padding:0}">
        <template #extra><a href="#">Edit</a></template>
        <div style="height: 130px;display: flex">
          <a-avatar :size="100" shape="square" style="margin: 15px;">
            <template #icon><UserOutlined /></template>
          </a-avatar>
          <div class="ellipsis-multiline2">
            &nbsp; &nbsp;  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently. We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.
          </div>
        </div>
        <template #actions>
          <span><HeartOutlined /> 999+</span>
          <LikeOutlined ></LikeOutlined>
          <DislikeOutlined />
        </template>
      </a-card>
      <a-card title="Default size card" style="width: 100%;text-align: left" :hoverable="true" :bodyStyle="{padding:0}">
        <template #extra><a href="#">Edit</a></template>
        <div style="height: 130px;display: flex">
          <a-avatar :size="100" shape="square" style="margin: 15px;">
            <template #icon><UserOutlined /></template>
          </a-avatar>
          <div class="ellipsis-multiline2">
            &nbsp; &nbsp;  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently. We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.
          </div>
        </div>
        <template #actions>
          <span><HeartOutlined /> 999+</span>
          <LikeOutlined ></LikeOutlined>
          <DislikeOutlined />
        </template>
      </a-card>
      <a-card title="Default size card" style="width: 100%;text-align: left" :hoverable="true" :bodyStyle="{padding:0}">
        <template #extra><a href="#">Edit</a></template>
        <div style="height: 130px;display: flex">
          <a-avatar :size="100" shape="square" style="margin: 15px;">
            <template #icon><UserOutlined /></template>
          </a-avatar>
          <div class="ellipsis-multiline2">
            &nbsp; &nbsp;  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently. We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.
          </div>
        </div>
        <template #actions>
          <span><HeartOutlined /> 999+</span>
          <LikeOutlined ></LikeOutlined>
          <DislikeOutlined />
        </template>
      </a-card>
      <a-card title="Default size card" style="width: 100%;text-align: left" :hoverable="true" :bodyStyle="{padding:0}">
        <template #extra><a href="#">Edit</a></template>
        <div style="height: 130px;display: flex">
          <a-avatar :size="100" shape="square" style="margin: 15px;">
            <template #icon><UserOutlined /></template>
          </a-avatar>
          <div class="ellipsis-multiline2">
            &nbsp; &nbsp;  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently. We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.
          </div>
        </div>
        <template #actions>
          <span><HeartOutlined /> 999+</span>
          <LikeOutlined ></LikeOutlined>
          <DislikeOutlined />
        </template>
      </a-card>
    </a-col>
    <a-col :md="5" :sm="24" :xs="24"  style="background-color: #1a1a1a" >
      3
    </a-col>
  </a-row>
</template>

<script setup>
import {UserOutlined, HeartOutlined, LikeOutlined, DislikeOutlined } from '@ant-design/icons-vue';
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
.ellipsis-multiline2{
  flex: 1; /* 右边部分自适应 */
  margin: 15px 15px 15px 0;
  /*word-break: break-all;*/
  /*word-wrap: break-word;*/
  /*overflow: hidden;*/
  /*text-overflow: ellipsis;*/
  /*display: -webkit-box;*/
  /*-webkit-box-orient: vertical;*/
  /*-webkit-line-clamp: 3;*/


  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 4; /* 显示3行 */
  overflow: hidden;
  font-size: 15px;
  text-overflow: ellipsis;
}
.ellipsis-multiline {
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 3;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
