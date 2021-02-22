<template>
  <header>
    <nav class="nav-extended">
      <div class="nav-wrapper">
        <router-link class="logo" to="/">Gaidai Chat Online</router-link>
        <a href="#" data-target="mobile-demo" class="sidenav-trigger">
          <i class="material-icons">menu</i>
        </a>
      </div>
    </nav>

    <ul class="sidenav" id="mobile-demo">
      <li>
        <router-link to="/posts/new">
          <i class="material-icons">create</i>Study Abroad
        </router-link>
      </li>
      <li>
        <router-link to="/articles/new">
          <i class="material-icons">create</i>Job Seeking
        </router-link>
      </li>
      <li>
        <router-link to="/questions/new">
          <i class="material-icons">create</i>Questions
        </router-link>
      </li>
    </ul>
  </header>
</template>

<script>
export default {
  mounted() {
    // jquery for sidenav
    $(document).ready(function() {
      $(".sidenav").sidenav();
    });
  },
  updated() {
    // jquery for sidenav
    $(document).ready(function() {
      $(".sidenav").sidenav();
    });
  }
};
</script>

<style lang="scss" scoped>
</style>