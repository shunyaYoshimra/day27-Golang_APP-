<template>
  <div id="question">
    <div class="card-panel">
      <span class="black-text">
        <router-link class="user-link" :to="{name: 'profile', params: {id: user.id}}">{{user.name}}</router-link>
        <p>{{question.content}}</p>
        <time>{{question.created_at}}</time>
        <template v-if="question.status">
          <span>解決済み</span>
        </template>
        <template v-else-if="!question.status && question.user_id === me.id">
          <button @click="finishQuestion()" class="btn grey lighten-1">Solved</button>
        </template>
      </span>
    </div>
    <div class="answers-wrapper">
      <small>回答(answers)</small>
      <hr />
      <div class="answers" v-for="(answer, id) in answers" :key="id">
        <div class="each-answer">
          <template v-if="editNum === id">
            <textarea v-model="answer.content"></textarea>
            <span @click="updateAnswer(answer.content, answer.id, question.id)">
              <i class="material-icons">edit</i>編集終了
            </span>
          </template>
          <template v-else>
            <p>{{answer.content}}</p>
            <span @click="editNum = id" v-if="answer.user_id === me.id && !question.status">
              <i class="material-icons">edit</i>編集
            </span>
          </template>
          <small>{{answer.created_at}}</small>
        </div>
      </div>
    </div> 
    <template v-if="!question.status">
      <div class="answer-form">
        <hr />
        <small>質問に回答してみましょう。(Answer this question)</small>
        <p>{{err}}</p>
        <textarea v-model="content"></textarea>
        <button @click="createAnswer()" class="btn grey lighten-1">answer</button>
      </div>
    </template>
  </div>
</template> 

<script>
import axios from "axios";
export default {
  data() {
    return {
      err: "",
      user: {},
      question: {},
      content: "",
      answers: [],
      editNum: "",
      me: {}
    };
  },
  mounted() {
    // get question
    axios.get(`/api/question/${this.$route.params.id}`).then(res => {
      this.question = res.data.data;
      // get asking user
      axios.get(`/api/user/${this.question.user_id}`).then(res => {
        this.user = res.data.data;
      });
    });
    // // get answers
    axios.get(`/api/answers/${this.$route.params.id}`).then(res => {
      console.log(res.data);
      for (let i = 0; i < res.data.length; i++) {
        this.answers.push(res.data[i]);
      }
    });
    // get me
    axios.get("/api/get_me").then(res => {
      this.me = res.data.data;
    });
  },
  methods: {
    createAnswer(content) {
      const params = new URLSearchParams();
      params.append("content", this.content);
      axios
        .post(`/api/answers/${this.question.id}`, params)
        .then(res => {
          this.answers.unshift(res.data.data);
        })
        .catch(err => {
          this.err = err.response.message;
        });
      this.content = "";
    },
    updateAnswer(editContent, answerID, questionID) {
      this.answers = [];
      const params = new URLSearchParams();
      params.append("content", editContent);
      params.append("answerID", answerID);
      params.append("questionID", questionID);
      axios.put("/api/answers", params).then(res => {
        for (let i = 0; i < res.data.length; i++) {
          this.answers.push(res.data[i]);
        }
      });
      this.editNum = false;
    },
    finishQuestion() {
      if (
        confirm(
          "回答の受付を終了しますか？(Are you sure to stop accepting answers?)"
        )
      ) {
        axios.put(`/api/change_status/${this.question.id}`);
        this.question.status = true;
      }
    }
  }
};
</script>

<style lang="scss" scoped>
#question {
  width: 600px;
  margin: 0 auto;

  .card-panel {
    .user-link {
      color: #333;
      font-weight: bold;
    }
    .user-link:hover {
      color: #999;
    }
    .btn {
      height: 2rem;
      font-size: 0.875rem;
      margin-left: 300px;
    }
  }

  .answers-wrapper {
    margin-top: 40px;
    margin-bottom: 250px;

    .answers {
      background-color: #fff;
      border: #999 solid 1px;
      margin-bottom: 10px;

      .each-answer {
        border-radius: 2px;
        padding: 10px;
        word-break: break-all;

        textarea {
          border: none;
        }

        textarea:focus {
          outline: none;
        }

        span {
          cursor: pointer;
        }

        small {
          margin-left: 1rem;
        }
      }
    }
  }

  .answer-form {
    position: fixed;
    padding-bottom: 30px;
    bottom: 0px;
    z-index: 10;
    width: 600px;
    background-color: #fafafa;

    textarea {
      border: 1px solid #666;
      border-radius: 5px;
      background-color: #fff;
      height: 80px;
    }

    textarea:focus {
      outline: none;
    }
  }
}
@media (max-width: 480px) {
  #question {
    width: 98%;
    margin: 20px auto;
    .card-panel {
      .btn {
        margin-left: 0px;
        display: block;
      }
    }
    .answer-form {
      width: 98%;
      bottom: 40px;
    }
  }
}
</style>