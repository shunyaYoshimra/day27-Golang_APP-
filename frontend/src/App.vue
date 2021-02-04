<template>
  <div id="app">
    <component :is="designByWidth"></component>
    <div class="main">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Header from './components/shared/Header.vue'
import MobileHeader from "./components/shared/MobileHeader.vue"
export default {
  components: {
    Header,
    MobileHeader
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
    }
  }
}
</script>

<style lang="scss">
.main {
  margin-top: 100px;
}
</style>
