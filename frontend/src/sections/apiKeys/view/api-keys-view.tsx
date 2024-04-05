import Container from "@mui/material/Container";

import { _jobs } from "src/_mock";

import { useSettingsContext } from "src/components/settings";

import { Typography } from "@mui/material";
import ApiKeysForm from "../api-keys-form";

export default function ApiKeysView() {
  const settings = useSettingsContext();

  return (
    <Container maxWidth={settings.themeStretch ? false : "lg"}>
      <Typography variant="h4" mb={5} >Api Keys</Typography>

      <ApiKeysForm />
    </Container>
  );
}
