import { Box } from "@mantine/core";
import { useForm } from "@mantine/form";
import { Checkbox, Button, Group, TextInput } from "@mantine/core";
import { useAuth } from "../../hooks/useAuth";
import { Navigate } from "react-router-dom";

export default function Login(props) {
  const { login } = useAuth();

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

    login({
      email: form.values.email,
      password: form.values.password
    });

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

          <TextInput
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
