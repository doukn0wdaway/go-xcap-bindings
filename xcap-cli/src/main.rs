use serde::Serialize;
use std::env;
use std::path::Path;

#[derive(Serialize)]
struct MonitorInfo {
    id: u32,
    name: String,
    width: u32,
    height: u32,
    frequency: f32,
    is_primary: bool,
    scale_factor: f32,
}

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
            let monitors = xcap::Monitor::all()?;
            let result: Vec<MonitorInfo> = monitors
                .into_iter()
                .map(|m| -> Result<MonitorInfo, xcap::XCapError> {
                    Ok(MonitorInfo {
                        id: m.id()?,
                        name: m.name()?,
                        width: m.width()?,
                        height: m.height()?,
                        frequency: m.frequency()?,
                        is_primary: m.is_primary()?,
                        scale_factor: m.scale_factor()?,
                    })
                })
                .collect::<Result<_, _>>()?;

            println!("{}", serde_json::to_string_pretty(&result)?);
        }

        "screenshot" => {
            let monitor_id: u32 = args
                .next()
                .ok_or("Missing monitorId")?
                .parse()
                .map_err(|_| "monitorId must be a number")?;
            let output = args.next().unwrap_or_else(|| "screenshot.png".to_string());

            let monitors = xcap::Monitor::all()?;
            let monitor = monitors
                .into_iter()
                .find(|m| m.id().unwrap_or(u32::MAX) == monitor_id)
                .ok_or(format!("Monitor with id {} not found", monitor_id))?;

            let image = monitor.capture_region(1600, 0, 500, 500)?;
            image.save(Path::new(&output))?;
            println!(
                "Monitor width: {}, height: {}",
                monitor.width()?,
                monitor.height()?
            );
            println!("Image width: {}, height: {}", image.width(), image.height());
            println!("Screenshot saved to {}", output);
        }

        _ => {
            eprintln!("Unknown command: {}", command);
            std::process::exit(1);
        }
    }

    Ok(())
}
