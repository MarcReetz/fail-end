import { Box, TextInput,Button } from "@mantine/core";
import { useForm } from "@mantine/form";
import Data from "../../const/const";

export default function AddFail() {
  const form = useForm({
    initialValues: {
      name: "",
      description: "",
    },
  });

  const onSubmit = () => {
    const response = fetch(Data.urls.apiFail, {
      method: "Post",
      credentials: 'include',
      body: JSON.stringify({
        title: form.values.name,
        description: form.values.description,
      }),
    });

    response.then( response => {
      if(response.status === 200){
        console.log("Got a Response")
      }
    })
  };

  return (
    <Box>
      <form onSubmit={form.onSubmit(onSubmit)}>
        <TextInput label="Name" {...form.getInputProps("name")} />
        <TextInput label="Description" {...form.getInputProps("description")} />
        <Button type="submit">Submit</Button>
      </form>
    </Box>
  );
}
