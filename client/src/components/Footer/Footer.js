import { Anchor, Group } from "@mantine/core";
import { createStyles } from "@mantine/core";
import { Link } from "react-router-dom";

const useStyles = createStyles((theme) => ({
  footer: {
    marginTop: 120,
    borderTop: `1px solid ${
      theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2]
    }`,
  },

  links: {
    [theme.fn.smallerThan("sm")]: {
      marginTop: theme.spacing.lg,
      marginBottom: theme.spacing.sm,
    },
  },

  inner: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: `${theme.spacing.md}px ${theme.spacing.md}px`,
    flexDirection: 'column',
  }
}));

const links = [{ name: "imprint", link: "/imprint" },{ name: "privacy", link:"/privacy"}];

export default function Footer() {
  var items = links.map((link) => {
    return (
      <Anchor component={Link} to={link.link}>
        {link.name}
      </Anchor>
    );
  });
  const { classes } = useStyles();
  return (
    <footer className={classes.footer}>
      <div className={classes.inner}>
        <Group className={classes.links}>{items}</Group>
      </div>
    </footer>
  );
}
