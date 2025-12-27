use screenshots::Screen;
use screenshots::display_info::DisplayInfo;
use std::env;
use std::path::Path;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut args = env::args();
    let _program = args.next(); // пропускаем имя программы

    let command = match args.next() {
        Some(cmd) => cmd,
        None => {
            eprintln!("Usage:");
            eprintln!("  get-monitors");
            eprintln!("  screenshot <monitorId> [output.png]");
            std::process::exit(1);
        }
    };

    match command.as_str() {
        "get-monitors" => {
            let monitors = Screen::all()?;
            let result: Vec<DisplayInfo> = monitors.iter().map(|m| m.display_info).collect();

            println!("{:#?}", result);
        }

        "screenshot" => {
            let monitor_id: u32 = args
                .next()
                .ok_or("Missing monitorId")?
                .parse()
                .map_err(|_| "monitorId must be a number")?;
            let output = args.next().unwrap_or_else(|| "screenshot.png".to_string());

            let monitors = Screen::all()?;
            let monitor = monitors
                .into_iter()
                .find(|m| m.display_info.id == monitor_id)
                .ok_or(format!("Monitor with id {} not found", monitor_id))?;

            let image = monitor.capture()?;
            image.save(Path::new(&output))?;
            println!("Screenshot saved to {}", output);
        }

        _ => {
            eprintln!("Unknown command: {}", command);
            std::process::exit(1);
        }
    }

    Ok(())
}
