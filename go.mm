<map version="0.9.0">
<!-- To view this file, download free mind mapping software FreeMind from http://freemind.sourceforge.net -->
<node CREATED="1292421082163" ID="ID_583505405" MODIFIED="1606209373698" TEXT="Go">
<node CREATED="1606209374203" ID="ID_820499233" MODIFIED="1606209385687" POSITION="right" TEXT="Project structure">
<node CREATED="1606210251718" ID="ID_1232302986" MODIFIED="1606210253128" TEXT="Modules">
<node CREATED="1606210253657" ID="ID_628930152" MODIFIED="1606210260408" TEXT="Starts from the dir with go.mod">
<node CREATED="1606210265638" ID="ID_993422522" MODIFIED="1606210292588" TEXT="defines module name">
<node CREATED="1606210272813" ID="ID_281360633" MODIFIED="1606210277646" TEXT="com.example/module"/>
</node>
<node CREATED="1606210293101" ID="ID_405778564" MODIFIED="1606210341945" TEXT="defines module deps">
<node CREATED="1606210299910" ID="ID_591305708" MODIFIED="1606210316928" TEXT="are installed with go install into $GOPATH"/>
<node CREATED="1606210344784" ID="ID_304147369" MODIFIED="1606210362951" TEXT="local deps can be overwritten with dir path">
<node CREATED="1606210363450" ID="ID_449706161" MODIFIED="1606210364136" TEXT="replace go-learn/other =&gt; ../other"/>
<node CREATED="1606210366304" ID="ID_1423988928" MODIFIED="1606210403098" TEXT="requires `go mod tidy` to load"/>
</node>
</node>
</node>
</node>
<node CREATED="1606210443724" ID="ID_94338017" MODIFIED="1606210445263" TEXT="Packages">
<node CREATED="1606210445927" ID="ID_1061341468" MODIFIED="1606210461323" TEXT="Defined in subdirectories of module">
<node CREATED="1606210462815" ID="ID_32923485" MODIFIED="1606210468792" TEXT="One subdirectory one package"/>
<node CREATED="1606210469141" ID="ID_1899416163" MODIFIED="1606210477944" TEXT="Package can be named differently than directory">
<node CREATED="1606210478484" ID="ID_1977454970" MODIFIED="1606210494672" TEXT="However, by convention better to leave dir name as package name"/>
</node>
</node>
<node CREATED="1606210517308" ID="ID_1588390655" MODIFIED="1606210520755" TEXT="Can be nested">
<node CREATED="1606210521146" ID="ID_148384739" MODIFIED="1606210524778" TEXT="Dir a/b/c"/>
<node CREATED="1606210525080" ID="ID_1845227532" MODIFIED="1606210528950" TEXT="Package a/b/c"/>
</node>
<node CREATED="1606210530176" ID="ID_143182019" MODIFIED="1606210572639" TEXT="Using packages from different modules uses module name as a first component in `import` clause">
<node CREATED="1606210574539" ID="ID_1573236806" MODIFIED="1606210579753" TEXT="Module X package a/b"/>
<node CREATED="1606210580122" ID="ID_1700619561" MODIFIED="1606210583617" TEXT="Import X/a/b"/>
</node>
<node CREATED="1606210585946" ID="ID_457772550" MODIFIED="1606210605548" TEXT="By default package exports only definitions started with capital letter"/>
</node>
</node>
</node>
</map>
