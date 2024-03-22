# Start-Job -ScriptBlock { npx tailwind -i .\styles.css -o .\public\styles.css --watch }
Start-Job -ScriptBlock { npx browser-sync start `
        --files 'views/**/*.html, public/**/*' `
        --port 3001 `
        --proxy 'localhost:3000' `
        --middleware 'function(req, res, next) { `
      res.setHeader("Cache-Control", "no-cache, no-store, must-revalidate"); `
      return next(); `
    }'
}
Start-Process pwsh "-Command", { npx tailwind -i .\styles.css -o .\public\styles.css --watch }
Start-Process pwsh "-Command", { npx browser-sync start `
        --files 'views/**/*.html, public/**/*' `
        --port 3001 `
        --proxy 'localhost:3000' `
        --middleware 'function(req, res, next) { res.setHeader("Cache-Control", "no-cache, no-store, must-revalidate"); return next(); }'
}
air