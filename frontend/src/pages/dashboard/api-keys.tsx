import { Helmet } from 'react-helmet-async';
import { ApiKeysView } from 'src/sections/apiKeys/view';


export default function ApiKeysPage() {
  return (
    <>
      <Helmet>
        <title> Dashboard: Api Keys</title>
      </Helmet>

      <ApiKeysView />
    </>
  );
}
