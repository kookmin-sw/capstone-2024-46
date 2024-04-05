import Box from "@mui/material/Box";
import Drawer from "@mui/material/Drawer";
import Stack from "@mui/material/Stack";

import { usePathname } from "src/routes/hooks";

import { useMockedUser } from "src/hooks/use-mocked-user";
import { useResponsive } from "src/hooks/use-responsive";

import Logo from "src/components/logo";
import { NavSectionVertical } from "src/components/nav-section";
import Scrollbar from "src/components/scrollbar";

import { NAV } from "../config-layout";
import { useNavData } from "./config-navigation";

export default function NavVertical() {
  const { user } = useMockedUser();

  const lgUp = useResponsive("up", "lg");

  const navData = useNavData();

  const renderContent = (
    <Scrollbar
      sx={{
        height: 1,
        "& .simplebar-content": {
          height: 1,
          display: "flex",
          flexDirection: "column",
        },
      }}
    >
      <Logo
        text="LOGO"
        sx={{ mt: 3, ml: 4, mb: 1, color: "black", fontWeight: "bold" }}
      />

      <NavSectionVertical
        data={navData}
        slotProps={{
          currentRole: user?.role,
        }}
      />

      <Box sx={{ flexGrow: 1 }} />
    </Scrollbar>
  );

  return (
    <Box
      sx={{
        flexShrink: { lg: 0 },
        width: { lg: NAV.W_VERTICAL },
      }}
    >
      <Stack
        sx={{
          height: 1,
          position: "fixed",
          width: NAV.W_VERTICAL,
          borderRight: (theme) => `dashed 1px ${theme.palette.divider}`,
        }}
      >
        {renderContent}
      </Stack>
    </Box>
  );
}
