<template>
  <div id="about-us">
    <div class="top-wrapper animate__animated animate__bounce">
      <h1 class="title">
        Welcome To
        <br />Gaidai Chat Online
      </h1>
      <p>This Application is for students of Kobe City University Of Foreign Studies</p>
    </div>
    <template v-if="visibleAbout">
      <div class="about-wrapper animate__animated animate__fadeIn">
        <h2 class="section-title">How To Use And What This Is For</h2>
        <p class="about">
          GaiDai Chat Onineは外大生の交流の促進、OBやOGとのコネクションの確立を目的としています。
          神戸市外国語大学がより魅力的な大学となるように、他の外大生に向けてあなたの経験を発信、共有してみましょう。
          また、他の人の投稿があなたのより良い留学の選択や就活の道標となるかもしれません。
        </p>
        <p class="about-en">
          Gaidai Chat Online is designed for the purpose of promoting communications among Gaidai students
          and establishing connections between enrolled students and alumnus or alumna.
          Let's share your experience to make our college more tempting.
          Also, Information posted by others would allow you to make better decisions on studying abroad or job hunting.
        </p>
      </div>
    </template>
    <template v-if="visiblePost">
      <div class="post-wrapper animate__animated animate__fadeIn">
        <h2 class="section-title">Share Your Study Abroad Experience</h2>
        <img class="post-img" src="/dist/img/public/post_sample.png" />
        <p class="post">
          留学の経験や様子を投稿することができます。投稿は写真、記述、タグを含みます。写真や記述であなたの経験がどのよなものか
          共有してみましょう。また、タグに留学している国や大学、その他の情報を付けることによって、留学先を考えている人の選ぶ材料となります。
        </p>
        <p class="post-en">
          You can share your study abroad experience. Posts include photos, descriptions and tags.
          Your photos and descriptions show how your experiece is like. It is useful, as well, for those who are going to study abroad to tag Information of
          the country or college you are studying abroad in when they decide which place they should go
        </p>
      </div>
    </template>
    <template v-if="visibleArticle">
      <div class="article-wrapper animate__animated animate__fadeIn">
        <h2 class="section-title">Share Your Job Seeking Experience</h2>
        <img class="article-img" src="/dist/img/public/article_sample.png" />
        <p class="article">
          就活の様子または仕事の様子を記事にしてみましょう。業種のカテゴリーをつけてあなたのインターンや面接などの就活の様子、また、
          卒業生は仕事の様子を記事して投稿することができます。記事には、URLを添付することができます。詳しく就活の様子を記して、同じ業種を目指す
          後輩の目標になりましょう。
        </p>
        <p class="article-en">
          Write articles about how you look for a job or what your job is like. You can add a category of job kind to your article and describe
          the experience of internship or interview etc. URL can be added to your article as well. Write your experience in detail and become a guide for
          people who will seek a job from now on.
        </p>
      </div>
    </template>
    <template v-if="visibleQuestion">
      <div class="question-wrapper animate__animated animate__fadeIn">
        <h2 class="section-title">Ask and Answer Questions About College Life</h2>
        <img class="question-img" src="/dist/img/public/question_sample.png" />
        <p class="question">
          授業やサークル、アルバイトなど大学生活に関わる分からないことを質問することができます。分からないことがあれば、
          他の外大生に助けてもらいましょう。また、自分が知っている疑問が投稿された際には、質問に答えてあげましょう。
        </p>
        <p class="question-en">
          You can ask questions related to your school life such as classes, club activities or part time job.
          Get other people's help. And if you find questions about which you know the answer, let's answer it.
        </p>
      </div>
    </template>
    <transition mode="out-in" name="fade">
      <div v-if="!visibleSpace" key="blank" class="space"></div>
      <div v-if="visibleSpace" key="start" class="space">Let's start!</div>
    </transition>
  </div>
</template>
  
<script>
import VAnimateCss from "animate.css";
export default {
  data() {
    return {
      visibleAbout: false,
      visiblePost: false,
      visibleArticle: false,
      visibleQuestion: false,
      visibleSpace: false
    };
  },
  created() {
    window.addEventListener("scroll", this.handleScrollAbout);
    window.addEventListener("scroll", this.handleScrollPost);
    window.addEventListener("scroll", this.handleScrollArticle);
    window.addEventListener("scroll", this.handleScrollQuestion);
    window.addEventListener("scroll", this.handleScrollSpace);
  },
  destroyed() {
    window.removeEventListener("scroll", this.handleScrollAbout);
    window.removeEventListener("scroll", this.handleScrollPost);
    window.removeEventListener("scroll", this.handleScrollArticle);
    window.removeEventListener("scroll", this.handleScrollQuestion);
    window.removeEventListener("scroll", this.handleScrollSpace);
  },
  methods: {
    handleScrollAbout() {
      if (!this.visibleAbout) {
        var top = this.$el.getBoundingClientRect().top;
        this.visibleAbout = top < -90;
      }
    },
    handleScrollPost() {
      if (!this.visiblePost) {
        var top = this.$el.getBoundingClientRect().top;
        this.visiblePost = top < -500;
      }
    },
    handleScrollArticle() {
      if (!this.visibleArticle) {
        var top = this.$el.getBoundingClientRect().top;
        this.visibleArticle = top < -1300;
      }
    },
    handleScrollQuestion() {
      if (!this.visibleQuestion) {
        var top = this.$el.getBoundingClientRect().top;
        this.visibleQuestion = top < -2000;
      }
    },
    handleScrollSpace() {
      if (!this.visibleSpace) {
        var top = this.$el.getBoundingClientRect().top;
        this.visibleSpace = top < -2500;
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.section-title {
  font-size: 2rem;
}
#about-us {
  width: 600px;
  margin: auto;
  .space {
    height: 600px;
    font-size: 3rem;
    text-align: center;
    padding-top: 200px;
  }
  .top-wrapper {
    padding-top: 100px;
    text-align: center;
    .title {
      font-size: 3rem;
    }
  }
  .about-wrapper,
  .post-wrapper,
  .article-wrapper,
  .question-wrapper {
    margin-top: 180px;
    text-align: center;
  }
  .about-wrapper {
    .about {
      margin-top: 40px;
      font-weight: bold;
    }
    .about-en {
      margin-top: 20px;
      font-weight: bold;
    }
  }
  .post-wrapper {
    .post-img {
      margin-top: 40px;
      width: 400px;
      height: auto;
    }
    .post {
      margin-top: 40px;
      font-weight: bold;
    }
    .post-en {
      margin-top: 20px;
      font-weight: bold;
    }
  }
  .article-wrapper {
    .article-img {
      margin-top: 40px;
      width: 400px;
      height: auto;
    }
    .article {
      margin-top: 40px;
      font-weight: bold;
    }
    .article-en {
      margin-top: 20px;
      font-weight: bold;
    }
  }
  .question-wrapper {
    margin-bottom: 60px;
    .question-img {
      margin-top: 40px;
      width: 400px;
      height: auto;
    }
    .question {
      margin-top: 40px;
      font-weight: bold;
    }
    .question-en {
      margin-top: 20px;
      font-weight: bold;
    }
  }
}
</style>