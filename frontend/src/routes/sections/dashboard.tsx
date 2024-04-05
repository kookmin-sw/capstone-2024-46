import { lazy, Suspense } from "react";
import { Outlet } from "react-router-dom";

import { AuthGuard } from "src/auth/guard";
import DashboardLayout from "src/layouts/dashboard";

import { LoadingScreen } from "src/components/loading-screen";

// ----------------------------------------------------------------------

// OVERVIEW
const IndexPage = lazy(() => import("src/pages/dashboard/app"));
const OverviewEcommercePage = lazy(
  () => import("src/pages/dashboard/ecommerce")
);
const OverviewAnalyticsPage = lazy(
  () => import("src/pages/dashboard/analytics")
);
const OverviewBankingPage = lazy(() => import("src/pages/dashboard/banking"));
const OverviewBookingPage = lazy(() => import("src/pages/dashboard/booking"));
const OverviewFilePage = lazy(() => import("src/pages/dashboard/file"));
// PRODUCT
const ProductDetailsPage = lazy(
  () => import("src/pages/dashboard/product/details")
);
const ProductListPage = lazy(() => import("src/pages/dashboard/product/list"));
const ProductCreatePage = lazy(() => import("src/pages/dashboard/product/new"));
const ProductEditPage = lazy(() => import("src/pages/dashboard/product/edit"));
// ORDER
const OrderListPage = lazy(() => import("src/pages/dashboard/order/list"));
const OrderDetailsPage = lazy(
  () => import("src/pages/dashboard/order/details")
);
// INVOICE
const InvoiceListPage = lazy(() => import("src/pages/dashboard/invoice/list"));
const InvoiceDetailsPage = lazy(
  () => import("src/pages/dashboard/invoice/details")
);
const InvoiceCreatePage = lazy(() => import("src/pages/dashboard/invoice/new"));
const InvoiceEditPage = lazy(() => import("src/pages/dashboard/invoice/edit"));
// USER
const UserProfilePage = lazy(() => import("src/pages/dashboard/user/profile"));
const UserCardsPage = lazy(() => import("src/pages/dashboard/user/cards"));
const UserListPage = lazy(() => import("src/pages/dashboard/user/list"));
const UserAccountPage = lazy(() => import("src/pages/dashboard/user/account"));
const UserCreatePage = lazy(() => import("src/pages/dashboard/user/new"));
const UserEditPage = lazy(() => import("src/pages/dashboard/user/edit"));
// BLOG
const BlogPostsPage = lazy(() => import("src/pages/dashboard/post/list"));
const BlogPostPage = lazy(() => import("src/pages/dashboard/post/details"));
const BlogNewPostPage = lazy(() => import("src/pages/dashboard/post/new"));
const BlogEditPostPage = lazy(() => import("src/pages/dashboard/post/edit"));
// JOB
const JobDetailsPage = lazy(() => import("src/pages/dashboard/job/details"));
const JobListPage = lazy(() => import("src/pages/dashboard/job/list"));
const JobCreatePage = lazy(() => import("src/pages/dashboard/job/new"));
const JobEditPage = lazy(() => import("src/pages/dashboard/job/edit"));
// TOUR
const TourDetailsPage = lazy(() => import("src/pages/dashboard/tour/details"));
const TourListPage = lazy(() => import("src/pages/dashboard/tour/list"));
const TourCreatePage = lazy(() => import("src/pages/dashboard/tour/new"));
const TourEditPage = lazy(() => import("src/pages/dashboard/tour/edit"));
// FILE MANAGER
const FileManagerPage = lazy(() => import("src/pages/dashboard/file-manager"));
const ApiKeysPage = lazy(() => import("src/pages/dashboard/api-keys"));
// APP
const ChatPage = lazy(() => import("src/pages/dashboard/chat"));
const MailPage = lazy(() => import("src/pages/dashboard/mail"));
const CalendarPage = lazy(() => import("src/pages/dashboard/calendar"));
const KanbanPage = lazy(() => import("src/pages/dashboard/kanban"));
// TEST RENDER PAGE BY ROLE
const PermissionDeniedPage = lazy(
  () => import("src/pages/dashboard/permission")
);
// BLANK PAGE
const BlankPage = lazy(() => import("src/pages/dashboard/blank"));

// ----------------------------------------------------------------------

export const dashboardRoutes = [
  {
    path: "dashboard",
    element: (
      <DashboardLayout>
        <Suspense fallback={<LoadingScreen />}>
          <Outlet />
        </Suspense>
      </DashboardLayout>
    ),
    children: [
      { element: <></>, index: true },

      { path: "knowledgebase", element: <FileManagerPage /> },
      { path: "apiKeys", element: <ApiKeysPage /> },
      { path: "settings", element: <></> },
    ],
  },
];
