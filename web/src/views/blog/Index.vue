<template>
  <a-row style="background-color: green;width: 100%"  >
    <a-col :md="5" :sm="24" :xs="24" style="background-color: #1a1a1a" >
      1
    </a-col>
    <a-col :md="14" :sm="24" :xs="24"  style="background-color: #992929;text-align: center" >
      <a-input-group compact  style="margin-top:30px;">
        <a-input v-model:value="value" placeholder="you can search blog title from here!"
                 style="width: calc(90% - 140px);height: 40px;background-color: transparent" />
        <a-button type="default" style="height: 40px;width: 70px" @click="onQuery">Query</a-button>
        <a-button type="default" style="height: 40px;width: 70px" @click="toRoute('/blog/edit/0')">New</a-button>
      </a-input-group>
      <a-card title="Default size card" class="blog-card" :hoverable="true" :bodyStyle="{padding:0}">
        <template #extra><a href="#">Edit</a></template>
        <div style="height: 130px;display: flex">
          <a-avatar :size="100"  style="margin: 15px;">
            <template #icon><UserOutlined /></template>
          </a-avatar>
          <div class="blog-content-eclipses">
           &nbsp; &nbsp;  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently. We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.
          </div>
        </div>
        <template #actions>
          <span><HeartOutlined /> 999+</span>
          <LikeOutlined ></LikeOutlined>
          <DislikeOutlined />
        </template>
      </a-card>
      <a-card title="Default size card" class="blog-card" :hoverable="true" :bodyStyle="{padding:0}">
        <template #extra><a href="#">Edit</a></template>
        <div style="height: 130px;display: flex">
          <a-avatar :size="100"  style="margin: 15px;">
            <a-image
                :width="100"
                src="https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png"
            />
          </a-avatar>

          <div class="blog-content-eclipses">
            &nbsp; &nbsp;  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently. We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.
          </div>
        </div>
        <template #actions>
          <span><HeartOutlined /> 999+</span>
          <LikeOutlined ></LikeOutlined>
          <DislikeOutlined />
        </template>
      </a-card>
      <a-card title="Default size card" class="blog-card" :hoverable="true" :bodyStyle="{padding:0}">
        <template #extra><a href="#">Edit</a></template>
        <div style="height: 130px;display: flex">
          <a-avatar :size="100"  style="margin: 15px;">
            <template #icon><UserOutlined /></template>
          </a-avatar>
          <div class="blog-content-eclipses">
            &nbsp; &nbsp;  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently. We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.
          </div>
        </div>
        <template #actions>
          <span><HeartOutlined /> 999+</span>
          <LikeOutlined ></LikeOutlined>
          <DislikeOutlined />
        </template>
      </a-card>
      <a-card title="Default size card" class="blog-card" :hoverable="true" :bodyStyle="{padding:0}">
        <template #extra><a href="#">Edit</a></template>
        <div style="height: 130px;display: flex">
          <a-avatar :size="100"  style="margin: 15px;">
            <template #icon><UserOutlined /></template>
          </a-avatar>
          <div class="blog-content-eclipses">
            &nbsp; &nbsp;  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.  We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently. We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.
          </div>
        </div>
        <template #actions>
          <span><HeartOutlined /> 999+</span>
          <LikeOutlined ></LikeOutlined>
          <DislikeOutlined />
        </template>
      </a-card>
      <a-card title="Default size card" class="blog-card" :hoverable="true" :bodyStyle="{padding:0}">
        <template #extra><a href="#">Edit</a></template>
        <div style="height: 130px;display: flex">
          <a-avatar :size="100"  style="margin: 15px;">
            <template #icon><UserOutlined /></template>
          </a-avatar>
          <div class="blog-content-eclipses">
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
import { useRouter } from 'vue-router';
const router = useRouter();
const toRoute=(path)=> {
  router.push(path)
}
</script>

<style scoped>
.blog-content-eclipses{
  flex: 1; /* 右边部分自适应 */
  margin: 15px 15px 15px 0;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 4; /* 显示3行 */
  overflow: hidden;
  font-size: 15px;
  text-overflow: ellipsis;
}

.blog-card{
  width: 90%;
  margin-top: 20px;
  text-align: left;
  margin-left: auto;
  margin-right: auto;
}
</style>
