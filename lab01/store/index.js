export const state = () => ({
  blogs: []
})

export const mutations = {
  ADD_BLOG(state, blog) {
    state.blogs.push(blog);
    console.log(blog);
  }
};
