export default function Topbar(): JSX.Element {
    return(
        <>
            <div className="p-6 max-w-full mx-auto bg-gray-200 shadow-lg flex items-center font-bebasneue">
                <div className="ml-auto">
                    <ul className="flex items-center space-x-4">
                        <li>
                            <a className="mr-3" href="/">Home</a>
                            <span className="border-r border border-slate-900"></span>
                        </li>
                        <li>
                            <a className="mr-3" href="#About">About</a>
                            <span className="border-r border border-slate-900"></span>
                        </li>
                        <li>
                            <a className="mr-3" href="#Portfolio">Portfolio</a>
                            <span className="border-r border border-slate-900"></span>
                        </li>
                        <li>
                            <a href="Contact">Contact</a>
                        </li>
                    </ul>
                </div>
            </div>
        </>
    )
}