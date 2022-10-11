import { ScrollArea, Table, createStyles, Button } from "@mantine/core";
import { useEffect, useState } from "react";
import Data from "./../../const/const";

const useStyles = createStyles((theme) => ({
  header: {
    "z-index": 2,
    position: "sticky",
    top: 0,
    backgroundColor:
      theme.colorScheme === "dark" ? theme.colors.dark[7] : theme.white,
    transition: "box-shadow 150ms ease",

    "&::after": {
      content: '""',
      position: "absolute",
      left: 0,
      right: 0,
      bottom: 0,
      borderBottom: `1px solid ${
        theme.colorScheme === "dark"
          ? theme.colors.dark[3]
          : theme.colors.gray[2]
      }`,
    },
  },

  scrolled: {
    boxShadow: theme.shadows.sm,
  },

  button: {
    "z-index": 0
  }
}));

export default function GetFailList() {
  //maybe set data to parent component
  const { classes, cx } = useStyles();
  const [scrolled, setScrolled] = useState(false);
  const [data, setData] = useState([]);

  useEffect(() => {
    const result = fetch(Data.urls.apiFail, {
      credentials: "include",
    });

    result.then((response) => {
      const json = response.json();
      json.then((value) => {
        setData(value);
        console.log(data);
      });
    });
  }, []);

  const addHit = (id) => {
    const result = fetch(Data.urls.apiFailHit1 + id + Data.urls.apiFailHit2, {
      credentials: "include",
      method: "Put",
    });

    result.then((response) => {
      if (response.ok) {
        const result = data.map(obj => {
          if (obj.id === id){
            return { ...obj, hits: obj.hits + 1}
          }
          return obj
        })
        setData(result)
      }
    });
  };

  const rows = data.map((row) => (
    <tr key={row.id}>
      <td>{row.title}</td>
      <td>{row.description}</td>
      <td>{row.hits}</td>
      <td>
        <Button className={classes.button} onClick={() => addHit(row.id)}>Hit</Button>
      </td>
    </tr>
  ));

  return (
    <>
      <ScrollArea
        sx={{ height: 500 }}
        onScrollPositionChange={({ y }) => setScrolled(y !== 0)}
      >
        <Table sx={{ minWidth: 700 }}>
          <thead
            className={cx(classes.header, { [classes.scrolled]: scrolled })}
          >
            <tr>
              <th>Fail</th>
              <th>Description</th>
              <th>hits</th>
              <th></th>
            </tr>
          </thead>
          <tbody>{rows}</tbody>
        </Table>
      </ScrollArea>
    </>
  );
}
