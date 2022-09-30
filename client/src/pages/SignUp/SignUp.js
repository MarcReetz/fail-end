import { Box, PasswordInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { Checkbox, Button, Group, TextInput } from "@mantine/core";
import { Navigate } from "react-router-dom";

export default function SignUp(props) {

  const form = useForm({
    initialValues: {
      email: "",
      password: "",
      confirmPassword: ""
    },

    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
      password: (value,values) => value !== values.confirmPassword  ? "Passwords did not match" : null,
      confirmPassword: (value,values) => value !== values.password ? "Passwords did not match" : null,
    },
  });

  const onSubmit = () => {


    if(props.next){
      return <Navigate to="/home" />;
    }else{
      return <Navigate to="/home" />;
    }
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

          <PasswordInput
            label="Confirm password"
            placeholder="***"
            {...form.getInputProps("confirmPassword")}
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
