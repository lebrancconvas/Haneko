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
            <Input v-model="blogForm.blogTitle" placeholder="Enter your blog title" />
          </FormItem>
          <FormItem label="Description">
            <Input v-model="blogForm.blogDescription" placeholder="Blog Description" />
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
            <Button type="success" @click="createBlog">Create Blog</Button>
          </FormItem>
        </Form>
      </div>
    </section>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        blogForm: {
          blogTitle: "",
          blogDescription: "",
          blogCategory: ""
        }
      }
    },
    methods: {
      createBlog() {
        const blogInformation = {
          title: this.blogForm.blogTitle,
          description: this.blogForm.blogDescription,
          category: this.blogForm.blogCategory
        }
        console.log(blogInformation);
        this.blogForm.blogTitle = "";
        this.blogForm.blogDescription = "";
        this.blogForm.blogCategory = "";
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
</style>
