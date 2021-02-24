<template>
  <div id="login">
    <div class="error">{{err}}</div>
    <div class="login-form card-panel">
      <p>LOG IN</p>
      <div class="inputs">
        <input type="text" v-model="email" placeholder="メール(Email)" />
        <input type="password" v-model="password" placeholder="パスワード(Password)" />
      </div>
      <button @click="login()" class="btn blue lighten-3">CLICK</button>
    </div>
    <div class="to-signup card-panel">
      <p>
        <router-link to="/signup">登録していない方はこちらへ(To Signup)</router-link>
      </p>
    </div>
    <div class="copy-right">
      <small>&copy; Shunya Yoshimura@2020</small>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      err: "",

      email: "",
      password: ""
    };
  },
  methods: {
    login() {
      const params = new URLSearchParams();
      params.append("email", this.email);
      params.append("password", this.password);
      axios
        .post("/api/login", params)
        .then(res => {
          console.log(res.data);
          this.$router.go("/");
        })
        .catch(err => {
          console.log(err.response);
          this.err = err.response.data.message;
        });
    }
  }
};
</script>

<style lang="scss" scoped>
#login {
  .error {
    text-align: center;
    color: tomato;
  }

  .login-form {
    width: 350px;
    height: 400px;
    margin: 0 auto;

    p {
      margin-top: 40px;
      color: #333;
      font-weight: bold;
      text-align: center;
    }

    .inputs {
      margin-top: 70px;

      input {
        border: 1px solid #666;
        border-radius: 10px;
        background-color: #fafafa;
        font-size: 0.875rem;
        color: #333;
        text-align: center;
      }

      input:focus {
        outline: none;
      }
    }

    .btn {
      margin-top: 20px;
      width: 300px;
    }
  }

  .to-signup {
    width: 350px;
    margin: 20px auto 0;
    height: 70px;

    p {
      transform: translateY(-15px);
      font-weight: bold;
    }
  }

  .copy-right {
    margin-top: 20px;
    text-align: center;

    small {
      color: #666;
    }
  }
}
@media (max-width: 480px) {
  #login {
    margin-bottom: 50px;
  }
}
</style>