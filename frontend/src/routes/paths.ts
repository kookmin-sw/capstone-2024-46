import { paramCase } from "src/utils/change-case";

import { _id, _postTitles } from "src/_mock/assets";

// ----------------------------------------------------------------------

const MOCK_ID = _id[1];

const MOCK_TITLE = _postTitles[2];

const ROOTS = {
  AUTH: "/auth",
  AUTH_DEMO: "/auth-demo",
  DASHBOARD: "/dashboard",
};

// ----------------------------------------------------------------------

export const paths = {
  comingSoon: "/coming-soon",
  maintenance: "/maintenance",
  pricing: "/pricing",
  payment: "/payment",
  about: "/about-us",
  contact: "/contact-us",
  faqs: "/faqs",
  page403: "/403",
  page404: "/404",
  page500: "/500",
  components: "/components",

  dashboard: {
    root: ROOTS.DASHBOARD,
    knowledgebase: `${ROOTS.DASHBOARD}/knowledgebase`,
    apiKeys: `${ROOTS.DASHBOARD}/apiKeys`,
    settings: `${ROOTS.DASHBOARD}/settings`,
  },
};
