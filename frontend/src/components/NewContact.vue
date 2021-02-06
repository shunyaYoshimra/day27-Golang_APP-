<template>
  <div id="new-contact">
    <div class="error">{{err}}</div>
    <div class="contact-area card-panel">
      <p>Contact</p>
      <small>このアプリは留学や就職について学生の交流を増やし大学の価値を高めることを目的としています。なるべく他のSNSとリンクさせメッセージのやり取りを行えるようにしましょう。</small>
      <!-- twitter -->
      <button
        class="waves-effect waves-light btn blue lighten-3 modal-trigger"
        href="#modal1"
      >Twitter</button>
      <div id="modal1" class="modal">
        <div class="modal-content">
          <p>Twitter</p>
          <small>@を抜いたアカウント名を入力してください(Please fill in your account name without @)</small>
          <input v-model="account" type="text" />
        </div>
        <div class="modal-footer">
          <a
            @click="createContact('twitter')"
            class="modal-close waves-effect waves-green btn-flat"
          >Next</a>
        </div>
      </div>
      <!-- instagram -->
      <button
        class="waves-effect waves-light btn pink lighten-3 modal-trigger"
        href="#modal2"
      >Instagram</button>
      <div id="modal2" class="modal">
        <div class="modal-content">
          <p>Instagram</p>
          <small>アカウント名を入力してください(Please fill in your account name)</small>
          <input v-model="account" type="text" />
        </div>
        <div class="modal-footer">
          <a
            @click="createContact('instagram')"
            class="modal-close waves-effect waves-green btn-flat"
          >Next</a>
        </div>
      </div>
      <p @click="createContact('')" class="later">後で登録する</p>
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
      account: ""
    };
  },
  mounted() {
    // for modals
    $(document).ready(function() {
      $(".modal").modal();
    });
  },
  methods: {
    createContact(media) {
      const params = new URLSearchParams();
      params.append("media", media);
      params.append("account", this.account);
      axios
        .post("/api/contacts", params)
        .then(res => {  
          console.log(res.data);
          this.$router.push("/start");
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
#new-contact {
  .error {
    text-align: center;
    color: tomato;
  }

  .contact-area {
    width: 350px;
    height: 400px;
    margin: 0 auto;

    p {
      margin-top: 40px;
      color: #333;
      font-weight: bold;
      text-align: center;
    }

    .modal {
      input {
        border: 1px solid #666;
        border-radius: 10px;
        background-color: #f5f5f5;
        font-size: 0.875rem;
        color: #333;
        text-align: center;
      }

      input:focus {
        outline: none;
      }
    }
  }

  .btn {
    margin-top: 20px;
    width: 300px;
  }

  .later {
    cursor: pointer;
  }

  .later:hover {
    color: #999;
  }

  .copy-right {
    margin-top: 20px;
    text-align: center;

    small {
      color: #666;
    }
  }
}
</style>