import Head from "next/head";
import Topbar from "./Topbar";
import Footer from "./Footer";


export default function Layout({children, home}: {children: any, home: boolean}) {
    return(
        <div>
            <Head>
                <link rel="icon" href="/favicon.ico" />
                <style>
                @import url('https://fonts.googleapis.com/css2?family=Bebas+Neue&family=Poppins:wght@100;200;300;400;500;600;700&display=swap');
                </style> 
            </Head>

            <header>
                <Topbar />
            </header>
            <main>{children}</main>
            <Footer />
        </div>
    )
}