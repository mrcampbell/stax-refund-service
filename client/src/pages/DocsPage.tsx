import SwaggerUI from "swagger-ui-react"
import "swagger-ui-react/swagger-ui.css"

function DocsPage() {
    return (
       <div className="bg-white">
         <SwaggerUI url="https://petstore.swagger.io/v2/swagger.json" />
       </div>
    );
}

export default DocsPage;