import { v4 as uuidv4 } from 'uuid';
import axios from 'axios';

export const state = () => ({
  blogs: []
})

export const mutations = {
  ADD_BLOG(state, blog) {
    const data = { id: uuidv4(), ...blog }
    state.blogs.push(data);
    axios.post("http://localhost:8000/api/v1/blogs", data)
      .then(res => {
        console.log(res);
      })
      .catch(err => {
        console.error(err);
        console.log(`[ERROR] Cannot Posting Data to API Server.`)
      })
  }
};
