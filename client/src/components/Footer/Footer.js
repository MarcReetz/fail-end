import { Anchor, Center, Group, Text } from "@mantine/core";
import { createStyles } from "@mantine/core";
import { Link } from "react-router-dom";

const useStyles = createStyles((theme) => ({

  link:{
    color: theme.colors.gray[5]
  },

  footer: {
    width: "100%",
    marginTop: 120,
    borderTop: `1px solid ${
      theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2]
    }`,
  },

  links: {
    marginTop: theme.spacing.lg,
    marginBottom: theme.spacing.sm,
  },

  inner: {
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center",
    padding: `${theme.spacing.md}px ${theme.spacing.md}px`,
    flexDirection: "column",

    [theme.fn.smallerThan("xs")]: {
      flexDirection: "row",
      justifyContent: "center",
      flexWrap: "nowrap",
    },
  },
}));

const links = [
  { name: "Imprint", link: "/imprint" },
  { name: "Privacy", link: "/privacy" },
];

export default function Footer() {
  const { classes } = useStyles();
  var items = links.map((link) => {
    return (
      <Anchor component={Link} to={link.link} key={link.name} className={classes.link}>
        {link.name}
      </Anchor>
    );
  });
  
  return (
    <footer className={classes.footer}>
      <div className={classes.inner}>
        <Group className={classes.links}>{items}</Group>
      </div>
      <Center>
        <Text color="dimmed" size="sm">
          © 2022
          <Anchor href="https://github.com/MarcReetz" target="_blank"> Marc Reetz </Anchor>
          All rights reserved.
        </Text>
      </Center>
    </footer>
  );
}
