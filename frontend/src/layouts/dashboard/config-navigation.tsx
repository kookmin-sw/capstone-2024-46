import { useMemo } from "react";

import { paths } from "src/routes/paths";

export function useNavData() {
  const data = useMemo(
    () => [
      {
        items: [
          {
            title: "HOME",
            path: paths.dashboard.root,
          },
          {
            title: "KNOWLEDGEBASE",
            path: paths.dashboard.knowledgebase,
          },
          {
            title: "API KEYs",
            path: paths.dashboard.apiKeys,
          },
          {
            title: "SETTINGS",
            path: paths.dashboard.settings,
          },
        ],
      },
    ],
    []
  );

  return data;
}
