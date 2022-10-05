import React from "react";
import Footer from "./Footer";
import { withRouter } from 'storybook-addon-react-router-v6';


export default {
  title:"Footer",
  component: Footer,
  decorators: [withRouter],
}

const Template = (args) => <Footer {...args}/>

export const First = Template.bind({})

First.args = {

}