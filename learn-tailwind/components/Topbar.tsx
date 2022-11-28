export default function Topbar(): JSX.Element {
    return(
        <>
            <div className="p-6 max-w-full mx-auto bg-gray-200 shadow-lg flex items-center space-x-4">
                <div className="ml-auto">
                    <ul className="flex items-center space-x-4">
                        <li>
                            <a href="/">Home</a>
                        </li>
                        <li>
                            <a href="#About">About</a>
                        </li>
                        <li>
                            <a href="#Portfolio">Portfolio</a>
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