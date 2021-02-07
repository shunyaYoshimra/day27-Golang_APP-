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

      <li class="dropdown-trigger" data-target="dropdown2">
        <i class="material-icons">add_circle_outline</i>
      </li>

      <!-- Dropdown Structure -->
      <ul id="dropdown2" class="dropdown-content">
        <li>
          <router-link class="test" to="/posts/new">Study Abroad</router-link>
        </li>
        <li>
          <router-link class="test" to="/articles/new">Job Seeking</router-link>
        </li>
        <li>
          <router-link class="test" to="/questions/new">Questions</router-link>
        </li>
      </ul>

      <li class="dropdown-trigger" data-target="dropdown3">
        <i class="material-icons">face</i>
      </li>

      <!-- Dropdown Structure -->
      <ul id="dropdown3" class="dropdown-content">
        <li>
          <router-link class="test" :to="{name: 'profile', params: {id: me.id}}">Profile</router-link>
        </li>
        <li>
          <a class="test" href="#!">two</a>
        </li>
        <li>
          <a class="test" href="#!">About Us</a>
        </li>
        <li class="divider" tabindex="-1"></li>
        <li>
          <a class="test" href="/api/logout">Logout</a>
        </li>
      </ul>
    </ul> 
  </header>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      me: {},
    }
  },
  mounted() {
    // jquery for dropdown
    $('.dropdown-trigger').dropdown();
    //  get me
    axios.get("/api/get_me").then((res) => {
      console.log(res.data)
      this.me = res.data.data;
    })
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