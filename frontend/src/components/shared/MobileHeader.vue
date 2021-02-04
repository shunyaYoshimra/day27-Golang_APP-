<template>
  <header>
    <nav class="nav-extended">
      <div class="nav-wrapper">
        <router-link class="logo" to="/">Gaidai Chat Online</router-link>
        <a href="#" data-target="mobile-demo" class="sidenav-trigger"><i class="material-icons">menu</i></a>
      </div>
    </nav>

    <ul class="sidenav" id="mobile-demo">
      <li><a href="sass.html">Sass</a></li>
      <li><a href="badges.html">Components</a></li>
      <li><a href="collapsible.html">JavaScript</a></li>
    </ul>
  </header>
</template>

<script>
export default {
  mounted() {
    // jquery for sidenav
    $(document).ready(function(){
      $('.sidenav').sidenav();
    });
  },
  updated() {
    // jquery for sidenav
    $(document).ready(function(){
      $('.sidenav').sidenav();
    });
  },
}
</script>

<style lang="scss" scoped>
.nav-wrapper {
  text-align: center;
}
</style>