<template>
  <div id="new-question">
    <div class="contents">
      <p>質問をする(Ask a question)</p>
      <div class="error">{{err}}</div>
      <input v-model="title" type="text" placeholder="タイトル(Title)" />
      <textarea v-model="content" placeholder="内容(Content)"></textarea>
      <select v-model="about">
        <option value disabled selected>About What?</option>
        <option value="classes">授業(classes)</option>
        <option value="seminer">ゼミ(seminer)</option>
        <option value="club activities">部活/サークル(club activities)</option>
        <option value="study abroad">留学(study abroad)</option>
        <option value="look for a job">就活(look for a job)</option>
        <option value="off capmus activities">校外活動(off capmus activities)</option>
        <option value="part time job">バイト(part time job)</option>
        <option value="other">その他(other)</option>
      </select>
      <label>About What?</label>
      <button @click="createQuestion()" class="btn grey lighten-1">Ask</button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      err: "",
      title: "",
      content: "",
      about: ""
    };
  },
  mounted() {
    // jquery for selecr
    $(document).ready(function() {
      $("select").formSelect();
    });
  },
  methods: {
    createQuestion() {
      const params = new URLSearchParams();
      params.append("title", this.title);
      params.append("content", this.content);
      params.append("about", this.about);
      axios
        .post("/api/questions", params)
        .then(res => {
          console.log(res.data);
          this.$router.push("/questions");
        })
        .catch(err => {
          this.err = err.response.data.message;
        });
    }
  }
};
</script>

<style lang="scss" scoped>
#new-question {
  width: 100%;
  height: 650px;
  background-color: #fff;

  .contents {
    margin: 0 auto;
    padding-top: 30px;
    width: 550px;

    .error {
      text-align: center;
      color: tomato;
    }

    p {
      font-weight: bold;
    }

    input {
      background-color: #fafafa;
      margin-top: 10px;
      border: 1px solid #333;
      font-size: 1rem;
      width: 100%;
      box-sizing: border-box;
      padding: 0 1rem;
    }

    input:focus {
      outline: none;
    }

    textarea {
      background-color: #fafafa;
      height: 300px;
      border: 1px solid #333;
      font-size: 1rem;
      padding: 1rem 1rem;
    }

    textarea:focus {
      outline: none;
    }
  }
}
@media (max-width: 480px) {
  #new-question {
    height: 550px;
    .contents {
      width: 98%;
      margin: auto;
    }
  }
}
</style>