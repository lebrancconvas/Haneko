<template>
  <div>
    <header>
      <div>
        <h1>Create Your Blog</h1>
      </div>
    </header>
    <section id="create-blog-section">
      <div>
        <Form>
          <FormItem label="Blog Title">
            <Input v-model="blogForm.blogTitle" placeholder="Enter your blog title" required />
          </FormItem>
          <FormItem label="Post (Max 1000 Alphabets)">
            <Input v-model="blogForm.blogDescription" size="large" maxlength="500" :rows="7" show-word-limit type="textarea" placeholder="Blog Description" required />
          </FormItem>
          <FormItem label="Category">
            <Select v-model="blogForm.blogCategory">
              <Option value="entertainment">Entertainment</Option>
              <Option value="knowledge">Knowledge</Option>
              <Option value="diary">Diary</Option>
            </Select>
          </FormItem>
          <FormItem>
            <Button type="error" @click="reset">Reset</Button>
            <Button type="primary" @click="createBlog">Post</Button>
            <div>
              <Modal v-model="modalSuccess" title="Success">
                <p> Post Blog Success. </p>
              </Modal>
              <Modal v-model="modalFail" title="Warning">
                <p>You must filled in all the required fields.</p>
              </Modal>
            </div>
          </FormItem>
        </Form>
      </div>
    </section>
  </div>
</template>

<script>
  // import axios from 'axios';
  export default {
    data() {
      return {
        blogForm: {
          blogTitle: "",
          blogDescription: "",
          blogCategory: ""
        },
        modalSuccess: false,
        modalFail: false
      }
    },
    methods: {
      createBlog() {
        if(this.blogForm.blogTitle === "" || this.blogForm.blogDescription === "" || this.blogForm.blogCategory === "") {
          this.modalFail = true;
          return;
        }
        const blogInformation = {
          title: this.blogForm.blogTitle,
          description: this.blogForm.blogDescription,
          category: this.blogForm.blogCategory
        }
        this.$store.commit('ADD_BLOG', blogInformation);
        this.blogForm.blogTitle = "";
        this.blogForm.blogDescription = "";
        this.blogForm.blogCategory = "";
        this.modalSuccess = true;
      },
      reset() {
        this.blogForm.blogTitle = "";
        this.blogForm.blogDescription = "";
        this.blogForm.blogCategory = "";
      }
    }
  }
</script>

<style scoped>
  * {
    text-align: center;
  }

  #create-blog-section {
    width: 50%;
    margin: 30px;
    margin-left: auto;
    margin-right: auto;
  }
</style>
