<template>
  <div id="app">
    <component :is="designByWidth"></component>
    <div class="main">
      <router-view></router-view>
    </div>
    <template v-if="this.width < 500">
      <MobileFooter></MobileFooter>
    </template>
  </div>
</template>

<script>
import Header from './components/shared/Header.vue';
import MobileHeader from "./components/shared/MobileHeader.vue";
import MobileFooter from "./components/shared/MobileFooter.vue";
export default {
  components: {
    Header,
    MobileHeader,
    MobileFooter,
  },
  data() {
    return {
      width: window.innerWidth,
    }
  },
  computed: {
    designByWidth: function() {
      self = this;
      if (self.width >= 500) {
        return Header
      } else {
        return  MobileHeader
      }
    },
  }
}
</script>

<style lang="scss">
#app {
  width: 100%;
  .main {
    margin-top: 100px;
  }
}
@media (max-width: 480px) {
  #app {
    .main{
      margin-top: 40px;
    }
  }
}
</style>
