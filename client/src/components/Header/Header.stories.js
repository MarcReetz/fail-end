import Header from "./Header";

export default {
  title: "header",
  component: Header
}

const Template = (args) => <Header {...args}/>

export const First = Template.bind({})

First.args = {
  user: {
    name: "marc"
  }
}