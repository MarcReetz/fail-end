import Footer from "../../components/Footer/Footer";
import { createStyles, Title, Text, Button, Container, Group } from '@mantine/core';
import { useNavigate } from "react-router-dom";

const useStyles = createStyles((theme) => ({
  root: {
    paddingTop: 80,
    paddingBottom: 80,
    height: "100vh"
  },

  label: {
    textAlign: 'center',
    fontWeight: 900,
    fontSize: 220,
    lineHeight: 1,
    marginBottom: theme.spacing.xl * 1.5,
    color: theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[2],

    [theme.fn.smallerThan('sm')]: {
      fontSize: 120,
    },
  },

  title: {
    fontFamily: `Greycliff CF, ${theme.fontFamily}`,
    textAlign: 'center',
    fontWeight: 900,
    fontSize: 38,

    [theme.fn.smallerThan('sm')]: {
      fontSize: 32,
    },
  },

  description: {
    maxWidth: 500,
    margin: 'auto',
    marginTop: theme.spacing.xl,
    marginBottom: theme.spacing.xl * 1.5,
  },
}));

export default function Page404 () {
  const {classes} = useStyles();
  const navigate = useNavigate();

  return (
    <>
    <Container className={classes.root}>
      <div className={classes.label}>404</div>
      <Title className={classes.title}>There is nothing here.</Title>
      <Text color="dimmed" size="lg" align="center" className={classes.description}>
        Unfortunately, this is only a 404 page. You may have mistyped the address, or the page has
        been moved to another URL.
      </Text>
      <Group position="center">
        <Button variant="subtle" size="md" onClick={() => navigate(`/`)}>
          Take me back to home page
        </Button>
        <Button variant="subtle" size="md" onClick={() => navigate(`/home/`)}>
          Take me to my profile
        </Button>
      </Group>
    </Container>
    <Footer/>
    </>
  )
}