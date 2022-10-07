import { ScrollArea, Table, createStyles} from "@mantine/core";
import { useState } from "react";
import Data from "./../../const/const"

const useStyles = createStyles((theme) => ({
  header: {
    position: 'sticky',
    top: 0,
    backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
    transition: 'box-shadow 150ms ease',

    '&::after': {
      content: '""',
      position: 'absolute',
      left: 0,
      right: 0,
      bottom: 0,
      borderBottom: `1px solid ${
        theme.colorScheme === 'dark' ? theme.colors.dark[3] : theme.colors.gray[2]
      }`,
    },
  },

  scrolled: {
    boxShadow: theme.shadows.sm,
  },
}));

export default function GetFailList() {

  const { classes, cx } = useStyles();
  const [scrolled, setScrolled] = useState(false);
  const [rows, setRows] = useState({});

  const result = fetch(Data.urls.apiFail)


  //PROBABly has to set to an action
  result.then( response => {
    const data = JSON.parse(response.json)

    const allRows = data.map((row) => (
      <tr key={row.title}>
        <td>{row.title}</td>
        <td>{row.description}</td>
        <td>{row.hits}</td>
      </tr>
    ));

    setRows(allRows)
  })

  //define row data 



  return (
    <>
      <ScrollArea
        sx={{ height: 300 }}
        onScrollPositionChange={({ y }) => setScrolled(y !== 0)}
      >
        <Table sx={{ minWidth: 700 }}>
          <thead
            className={cx(classes.header, { [classes.scrolled]: scrolled })}
          >
            <tr>
              <th>Name</th>
              <th>Email</th>
              <th>Company</th>
            </tr>
          </thead>
          <tbody>{rows}</tbody>
        </Table>
      </ScrollArea>
    </>
  );
}
