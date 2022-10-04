import { Navigate, Outlet } from "react-router-dom";
import Cookies from "js-cookie";
import { Center, Tabs } from "@mantine/core";
import { useNavigate, useParams } from "react-router-dom";
import { createStyles } from "@mantine/core";
import Footer from "../../components/Footer/Footer";

const useStyles = createStyles((theme) => ({
  header: {
    paddingTop: theme.spacing.sm,
    backgroundColor:
      theme.colorScheme === "dark"
        ? theme.colors.dark[6]
        : theme.colors.gray[0],
    borderBottom: `1px solid ${
      theme.colorScheme === "dark" ? "transparent" : theme.colors.gray[2]
    }`,
    marginBottom: 120,
    paddingBottom: 0,
  },

  tabs: {
    [theme.fn.smallerThan("sm")]: {
      display: "none",
    },
  },

  tabsList: {
    borderBottom: "0 !important",
  },

  tab: {
    fontWeight: 500,
    height: 38,
    backgroundColor: "transparent",

    "&:hover": {
      backgroundColor:
        theme.colorScheme === "dark"
          ? theme.colors.dark[5]
          : theme.colors.gray[1],
    },

    "&[data-active]": {
      backgroundColor:
        theme.colorScheme === "dark" ? theme.colors.dark[7] : theme.white,
      borderColor:
        theme.colorScheme === "dark"
          ? theme.colors.dark[7]
          : theme.colors.gray[2],
    },
  },
}));

export default function Home() {
  const navigate = useNavigate();
  const { tabValue } = useParams();

  const { classes, theme, cx } = useStyles();

  var allCookies = Cookies.get();
  console.log(allCookies);

  var user = Cookies.get("token");

  console.log(user);

  if (!user) {
    return <Navigate to="/login" />;
  }

  return (
    <>
      <header className={classes.header}>
        <h1>Home</h1>
        <nav>
          <Center>
            <Tabs
              variant="outline"
              classNames={{
                root: classes.tabs,
                tabsList: classes.tabsList,
                tab: classes.tab,
              }}
              value={tabValue}
              onTabChange={(value) => navigate(`/home/${value}`)}
            >
              <Tabs.List>
                <Tabs.Tab value="list" key="list">
                  fail
                </Tabs.Tab>
                <Tabs.Tab value="add" key="add">
                  Second tab
                </Tabs.Tab>
              </Tabs.List>
            </Tabs>
          </Center>
        </nav>
      </header>
      <Outlet />
      <Footer/>
    </>
  );
}
