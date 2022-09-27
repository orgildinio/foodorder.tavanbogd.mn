const path = require("path");
const mix = require('laravel-mix');
const fs = require('fs');
if (mix.inProduction()) {

    const fileName = './public/mix-manifest.json';
    fs.writeFile(fileName, "{}", function writeJSON(err) {});
    if (fs.existsSync('./public/assets/lambda/js')) {
        fs.rmSync('./public/assets/lambda/js', { recursive: true });
    }
}

require("laravel-mix-merge-manifest");

const lambdaRoot =  process.env.LAMBDA_ROOT;
const lambdaModule = `${lambdaRoot}/src/modules`;
const lambdaPlatform = `${lambdaRoot}/src`;
const dataform_custom = "assets/dataform_custom";
const datagrid_custom = "assets/datagrid_custom";

if (!fs.existsSync(dataform_custom)) {
    fs.mkdirSync(dataform_custom);
}

if (!fs.existsSync(datagrid_custom)) {
    fs.mkdirSync(datagrid_custom);
}

mix.webpackConfig({
    output: {
        chunkFilename: mix.inProduction()
            ? "assets/lambda/js/chunks/[name].[chunkhash].js"
            : "assets/lambda/js/chunks/[name].js"
    },
    resolve: {
        modules: [
            path.resolve(__dirname, './node_modules'),
            path.resolve(`${lambdaRoot}`, 'node_modules'),
        ],
        alias: {
            dataform_custom: path.resolve(__dirname, dataform_custom),
            datagrid_custom: path.resolve(__dirname, datagrid_custom),
            lambda_platform: path.resolve(__dirname, lambdaPlatform),
            vue$: "vue/dist/vue.common.js",
        },
        extensions: ["*", ".js", ".ts", ".vue", ".json"],
        symlinks: false
    }
});

mix.options({
    processCssUrls: false
}).setPublicPath("public");


//Vendor styles
mix.styles(
    [
        `node_modules/` + "iview/dist/styles/iview.css",
        `node_modules/` + "leaflet/dist/leaflet.css",
        `node_modules/` + "leaflet-draw/dist/leaflet.draw.css",
        `node_modules/` + "leaflet.fullscreen/Control.FullScreen.css",
        `node_modules/` + "vue-multiselect/dist/vue-multiselect.min.css",
        `node_modules/` + "ag-grid-community/dist/styles/ag-grid.css",
        `node_modules/` + "ag-grid-community/dist/styles/ag-theme-balham.css",
    ],
    "public/assets/lambda/css/vendor.css"
);

mix.extract(
    ["ag-grid-community", "ag-grid-enterprise", "ag-grid-vue"],
    "public/assets/lambda/js/datagrid-vendor.js"
);

mix.extract(
    [
        "sql-formatter",
        "sql-parser",
        "jQuery-QueryBuilder",
        "axios",
        "iview",
        "leaflet",
        "esri-leaflet",
        "leaflet.fullscreen",
        "leaflet-draw",
        "leaflet-rotatedmarker",
        "leaflet.gridlayer.googlemutant",
        "lodash",
        "normalize.css",
        "vue",
        "vuex",
        "vue-router"
    ],
    "public/assets/lambda/js/vendor.js"
)
    .version()
    .mergeManifest();

const compileModules = () => {

    //
    // mix.js("assets/admin/js/index.js", "public/assets/admin/js/app.js").vue();
    // mix.sass("assets/admin/scss/style.scss", "public/assets/admin/css/app.css");
    //
    // mix.js("assets/moqup/moqup_view.js", "public/assets/admin/js/moqup_view.js").vue();
    // mix.sass("assets/moqup/scss/moqup_view.scss", "public/assets/admin/css/moqup_view.css");

    mix .js(`${lambdaRoot}/src/index.js`, `public/assets/lambda/js/puzzle.js`).vue()
        .sass(`${lambdaRoot}/src/scss/style.scss`, `public/assets/lambda/css/puzzle.css`)
        .js(`${lambdaModule}/agent/index.js`, `public/assets/lambda/js/agent.js`).vue()
        .sass(`${lambdaModule}/agent/scss/agent.scss`, `public/assets/lambda/css/agent.css`)
        //
        // .js(`${lambdaModule}/agent/auth.js`, `public/assets/lambda/js/auth.js`).vue()
        // .sass(`${lambdaModule}/agent/scss/auth.scss`, `public/assets/lambda/css/auth.css`)

        .js(`${lambdaModule}/krud/index.js`, "public/assets/lambda/js/krud.js").vue()
        .sass(`${lambdaModule}/krud/scss/krud.scss`, `public/assets/lambda/css/krud.css`)

        .js(`${lambdaModule}/datagrid/index.js`, `public/assets/lambda/js/datagrid.js`).vue()
        .js(`${lambdaModule}/datagrid/index-builder.js`, `public/assets/lambda/js/datagrid-builder.js`).vue()
        .sass(`${lambdaModule}/datagrid/scss/style.scss`, `public/assets/lambda/css/datagrid.css`)

        .js(`${lambdaModule}/dataform/index.js`, `public/assets/lambda/js/dataform.js`).vue()
        .js(`${lambdaModule}/dataform/index-builder.js`, `public/assets/lambda/js/dataform-builder.js`).vue()
        .sass(`${lambdaModule}/dataform/scss/style.scss`, `public/assets/lambda/css/dataform.css`)

        .js(`${lambdaModule}/datasource/index.js`, `public/assets/lambda/js/datasource.js`).vue()
        .sass(`${lambdaModule}/datasource/scss/style.scss`, `public/assets/lambda/css/datasource.css`)

        // .js(`${lambdaModule}/page/index.js`, "public/assets/lambda/js/page.js").vue()
        // .js(`${lambdaModule}/page/page-nav.js`, "public/assets/lambda/js/page-nav.js").vue()
        //
        // .js(`${lambdaModule}/moqup/index.js`, `public/assets/lambda/js/moqup.js`).vue()
        // .sass(`${lambdaModule}/moqup/scss/style.scss`, `public/assets/lambda/css/moqup.css`)
        //
        // .js(`${lambdaModule}/chart/index.js`, "public/assets/lambda/js/chart.js").vue()
        // .js(`${lambdaModule}/chart/index-builder.js`, `public/assets/lambda/js/chart-builder.js`).vue()
        // .sass(`${lambdaModule}/chart/scss/style.scss`, "public/assets/lambda/css/chart.css")

        // .js(`${lambdaModule}/notify/index.js`, `public/assets/lambda/js/notification.js`).vue()

        // .js(`${lambdaModule}/logger/index.js`, "public/assets/lambda/logger.js")
        // .sass(`${lambdaModule}/logger/scss/logger.scss`, "public/assets/lambda/logger.css")
        // .sourceMaps()
};

const compileTemplate = () => {
    mix.js(`${lambdaRoot}/src/template/paper/index.js`, `public/assets/lambda/js/paper.js`).vue()
        .sass(`${lambdaRoot}/src//template/paper/scss/style.scss`, `public/assets/lambda/css/paper.css`);
};

console.log("COMPILING LAMBDA");
compileModules();
compileTemplate();
