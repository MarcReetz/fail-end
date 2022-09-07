import React from "react";
import Login from "./Login";

export default {
  title:"Login Page",
  component: Login
}

const Template = (args) => <Login {...args}/>

export const First = Template.bind({})

First.args = {

}