<template>
  <header class="header">
    <div>
      <router-link to="/" class="logo" >Gaidai Chat Online</router-link>
    </div>
    <ul class="header-nav">
      <li class="dropdown-trigger" data-target="dropdown1">
        <i class="material-icons">photo</i>
      </li>

      <!-- Dropdown Structure -->
      <ul id="dropdown1" class="dropdown-content">
        <li class="t">
          <router-link class="test" to="/">Study Abroad</router-link>
        </li>
        <li class="t">
          <router-link class="test" to="/">Job Seeking</router-link>
        </li>
        <li class="t">
          <router-link class="test" to="/q">Questions</router-link>
        </li>
      </ul>
    </ul> 
  </header>
</template>

<script>
export default {
  mounted() {
     // jquery for dropdown
    $('.dropdown-trigger').dropdown();
  },
  updated() {
     // jquery for dropdown
    $('.dropdown-trigger').dropdown();
  },
}
</script>

<style lang="scss">
.header {
  position: fixed;
  top: 0;
  z-index: 10;
  width: 100%;
  background-color: #fff;
  border-bottom: 1px solid #bdbdbd;
  display: flex;
  justify-content: space-around;
  align-items: center;

  .logo {
    font-weight: bold;
    color: #333;
  }

  .header-nav {
    position: relative;
    display: flex;
    list-style: none;

    li {
      margin-left: 15px;
      font-weight: bold;
    }

    .add-nav {
      position: absolute;
      width: 200px;
      top: 20px;
      right: 0;

      .add-link {
        color: #333;
        span {
          font-size: 1rem;
          margin-left: 1rem;
        }
      }

      .add-link:hover {
        color: #64b5f6;
      }
    }
  }
}
</style>