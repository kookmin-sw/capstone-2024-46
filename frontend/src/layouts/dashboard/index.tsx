import Box from "@mui/material/Box";

import Main from "./main";
import NavVertical from "./nav-vertical";

type Props = {
  children: React.ReactNode;
};

export default function DashboardLayout({ children }: Props) {
  return (
    <>
      <Box
        sx={{
          minHeight: 1,
          display: "flex",
          flexDirection: { xs: "column", lg: "row" },
        }}
      >
        <NavVertical />

        <Main>{children}</Main>
      </Box>
    </>
  );
}
