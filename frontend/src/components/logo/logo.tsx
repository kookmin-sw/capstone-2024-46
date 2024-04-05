import { forwardRef } from "react";

import Box, { BoxProps } from "@mui/material/Box";
import Link from "@mui/material/Link";

import { RouterLink } from "src/routes/components";

export interface LogoProps extends BoxProps {
  text?: string;
  disabledLink?: boolean;
}

const Logo = forwardRef<HTMLDivElement, LogoProps>(
  ({ text, disabledLink = false, sx, ...other }, ref) => {
    const logo = (
      <Box
        ref={ref}
        component="div"
        sx={{
          display: "inline-flex",
          ...sx,
        }}
        {...other}
      >
        {text}
      </Box>
    );

    if (disabledLink) {
      return logo;
    }

    return (
      <Link component={RouterLink} href="/" sx={{ display: "contents" }}>
        {logo}
      </Link>
    );
  }
);

export default Logo;
