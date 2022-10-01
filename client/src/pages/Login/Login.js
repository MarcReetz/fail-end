import { Box, PasswordInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { useNavigate } from "react-router-dom";
import { Checkbox, Button, Group, TextInput } from "@mantine/core";
import Data from "../../const/const"
import Cookies from "js-cookie";

export default function Login(props) {
  const navigate = useNavigate()

  const form = useForm({
    initialValues: {
      email: "",
      password: "",
    },

    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
    },
  });

  const onSubmit = () => {

    const response = fetch(Data.urls.getApiLogin,{
      credentials: "include",
      method: "POST",
      body: JSON.stringify({username: form.values.email,password: form.values.password})
    })

    response.then( (data) =>{
      console.log(data)
      var allCookies = Cookies.get()
      allCookies = document.cookie;
      console.log(allCookies)
      console.log(allCookies)
      if(props.next){
        console.log("first")
        navigate(props.next)
      }else{
        console.log("second")
        navigate("/home")
      }
    }
    )
    /*login({
      email: form.values.email,
      password: form.values.password
    });*/
  }

  return (
    <>
      <Box>
        <form onSubmit={form.onSubmit(onSubmit)}>
          <TextInput
            withAsterisk
            label="Email"
            placeholder="your@email.com"
            {...form.getInputProps("email")}
          />

          <PasswordInput
            label="Password"
            placeholder="***"
            {...form.getInputProps("password")}
          />
          <Group position="apart" mt="md">
            <Checkbox mt="md" label="Keep me Loged in" />
            <Button type="submit">Submit</Button>
          </Group>
        </form>
      </Box>
    </>
  );
}
