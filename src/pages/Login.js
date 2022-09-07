import { Box } from "@mantine/core";
import { useForm } from "@mantine/form";
import { Checkbox, Button, Group, TextInput } from "@mantine/core";

export default function Login() {
  const form = useForm({
    initialValues: {
      email: "",
      password: "",
    },

    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
    },
  });

  return (
    <>
      <Box>
        <form onSubmit={form.onSubmit((values) => console.log(values))}>
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

          <Group position="right" mt="md">
            <Checkbox mt="md" label="Keep me Loged in" />
            <Button type="submit">Submit</Button>
          </Group>
        </form>
      </Box>
    </>
  );
}
