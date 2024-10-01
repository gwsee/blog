<template>
  <a-row style="background-color: green;width: 100%"  >
    <a-col :md="5" :sm="24" :xs="24" style="background-color: #1a1a1a" >
      1
    </a-col>
    <a-col :md="14" :sm="24" :xs="24"  style="background-color: #992929" >
      <a-card hoverable style="width: 300px">
        <template #actions>
          <!--            <setting-outlined key="setting" />-->
          <!--            <edit-outlined key="edit" />-->
          <!--            <ellipsis-outlined key="ellipsis" />-->
        </template>
        <template #cover>
          <img
              alt="example"
              src="https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"
          />
        </template>
        <a-card-meta title="Card title" description="This is the description">
          <template #avatar>
            <a-avatar src="https://joeschmoe.io/api/v1/random" />
          </template>
        </a-card-meta>
      </a-card>
      <div style="background: #ececec; padding: 30px">
        <a-card title="Card title" :bordered="false" style="width: 300px">
          <p>Card content</p>
          <p>Card content</p>
          <p>Card content</p>
        </a-card>
      </div>
      <a-card title="Default size card" style="width: 300px">
        <template #extra><a href="#">more</a></template>
        <p>card content</p>
        <p>card content</p>
        <p>card content</p>
      </a-card>
      <a-card size="small" title="Small size card" style="width: 300px">
        <template #extra><a href="#">more</a></template>
        <p>card content</p>
        <p>card content</p>
        <p>card content</p>
      </a-card>
      <a-card hoverable style="width: 300px">
        <template #actions>
          <!--            <setting-outlined key="setting" />-->
          <!--            <edit-outlined key="edit" />-->
          <!--            <ellipsis-outlined key="ellipsis" />-->
        </template>
        <template #cover>
          <img
              alt="example"
              src="https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"
          />
        </template>
        <a-card-meta title="Card title" description="This is the description">
          <template #avatar>
            <a-avatar src="https://joeschmoe.io/api/v1/random" />
          </template>
        </a-card-meta>
      </a-card>
    </a-col>
    <a-col :md="5" :sm="24" :xs="24"  style="background-color: #1a1a1a" >
      3
    </a-col>
  </a-row>
</template>

<script>
export default {
}
</script>

<style scoped>

</style>
