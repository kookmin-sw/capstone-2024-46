import { yupResolver } from "@hookform/resolvers/yup";
import { useMemo } from "react";
import { Controller, useForm } from "react-hook-form";
import * as Yup from "yup";

import LoadingButton from "@mui/lab/LoadingButton";
import Box from "@mui/material/Box";
import ButtonBase from "@mui/material/ButtonBase";
import Card from "@mui/material/Card";
import FormControlLabel from "@mui/material/FormControlLabel";
import Paper from "@mui/material/Paper";
import Stack from "@mui/material/Stack";
import Switch from "@mui/material/Switch";
import Typography from "@mui/material/Typography";
import Grid from "@mui/material/Unstable_Grid2";

import { useRouter } from "src/routes/hooks";

import { JOB_BENEFIT_OPTIONS } from "src/_mock";

import FormProvider, {
  RHFRadioGroup,
  RHFTextField,
} from "src/components/hook-form";
import Iconify from "src/components/iconify";

export default function JobNewEditForm() {
  const router = useRouter();

  const NewApiKeysSchema = Yup.object().shape({
    type: Yup.string().required("Api type is required"),
    key: Yup.string().required("Api key is required"),
  });

  const defaultValues = useMemo(
    () => ({
      type: "Hourly",
      key: "",
    }),
    []
  );

  const methods = useForm({
    resolver: yupResolver(NewApiKeysSchema),
    defaultValues,
  });

  const {
    reset,
    control,
    handleSubmit,
    formState: { isSubmitting },
  } = methods;

  // useEffect(() => {
  //   if (currentApiKeys) {
  //     reset(defaultValues);
  //   }
  // }, [currentApiKeys, defaultValues, reset]);

  const onSubmit = handleSubmit(async (data) => {
    try {
      await new Promise((resolve) => setTimeout(resolve, 500));
      reset();
      // enqueueSnackbar(currentApiKeys ? "Update success!" : "Create success!");
      console.info("DATA", data);
    } catch (error) {
      console.error(error);
    }
  });

  const renderProperties = (
    <>
      <Grid xs={12} md={12}>
        <Card>
          <Stack spacing={3} sx={{ p: 3 }}>
            <Stack spacing={2}>
              <Typography variant="subtitle2">
                API KEY를 입력해주세요
              </Typography>

              <Controller
                name="type"
                control={control}
                render={({ field }) => (
                  <Box
                    gap={2}
                    display="grid"
                    gridTemplateColumns="repeat(3, 1fr)"
                  >
                    {[
                      {
                        label: "OPEN_AI",
                        icon: (
                          <Iconify icon="solar:clock-circle-bold" width={32} />
                        ),
                      },
                      {
                        label: "CLAUDE",
                        icon: (
                          <Iconify icon="solar:wad-of-money-bold" width={32} />
                        ),
                      },
                      {
                        label: "LLAMA2",
                        icon: (
                          <Iconify icon="solar:wad-of-money-bold" width={32} />
                        ),
                      },
                    ].map((item) => (
                      <Paper
                        component={ButtonBase}
                        variant="outlined"
                        key={item.label}
                        onClick={() => field.onChange(item.label)}
                        sx={{
                          p: 2.5,
                          borderRadius: 1,
                          typography: "subtitle2",
                          flexDirection: "column",
                          ...(item.label === field.value && {
                            borderWidth: 2,
                            borderColor: "text.primary",
                          }),
                        }}
                      >
                        {item.icon}
                        {item.label}
                      </Paper>
                    ))}
                  </Box>
                )}
              />

              <RHFTextField
                name="key"
                placeholder="Api Key"
                type="string"
                // InputProps={{
                //   startAdornment: (
                //     <InputAdornment position="start">
                //       <Box sx={{ typography: 'subtitle2', color: 'text.disabled' }}>$</Box>
                //     </InputAdornment>
                //   ),
                // }}
              />
            </Stack>
          </Stack>
        </Card>
      </Grid>
    </>
  );

  const renderActions = (
    <>
      <Grid
        xs={12}
        md={12}
        sx={{
          display: "flex",
          alignItems: "center",
          justifyContent: "flex-end",
        }}
      >
        <LoadingButton
          type="submit"
          variant="contained"
          size="large"
          loading={isSubmitting}
          sx={{ ml: 2 }}
        >
          Save
        </LoadingButton>
      </Grid>
    </>
  );

  return (
    <FormProvider methods={methods} onSubmit={onSubmit}>
      <Grid container spacing={3}>
        {renderProperties}

        {renderActions}
      </Grid>
    </FormProvider>
  );
}
